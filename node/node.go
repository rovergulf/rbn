package node

import (
	"context"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/rovergulf/rbn/core"
	"github.com/rovergulf/rbn/core/types"
	"github.com/rovergulf/rbn/database/badgerdb"
	"github.com/rovergulf/rbn/params"
	"github.com/rovergulf/rbn/pkg/sigutil"
	"github.com/rovergulf/rbn/pkg/traceutil"
	"github.com/rovergulf/rbn/wallets"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/raft/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"
	"sync"
	"time"
)

const (
	DbFileName = "node.db"

	DefaultNodeIP   = "127.0.0.1"
	DefaultNodePort = 9420
	HttpSSLPort     = 443

	endpointStatus  = "/node/status"
	endpointSync    = "/node/sync"
	endpointAddPeer = "/node/peer"
)

// Node represents blockchain network peer node
type Node struct {
	metadata PeerNode
	account  *wallets.Wallet

	config params.Options

	bc *core.Blockchain
	wm *wallets.Manager
	db *badger.DB

	grpcServer  *grpc.Server
	httpHandler httpServer

	inGenRace bool

	knownPeers knownPeers

	pendingTXs map[common.Hash]types.SignedTx

	newSyncBlocks chan types.Block
	newSyncTXs    chan types.SignedTx

	raftStorage *raft.MemoryStorage
	//Lock *sync.RWMutex

	logger *zap.SugaredLogger
	tracer traceutil.Tracer
}

// New creates and returns new node if blockchain available
func New(opts params.Options) (*Node, error) {
	nodeAddr := viper.GetString("node.addr")
	nodePort := viper.GetUint64("node.port")

	syncMode := viper.GetString("node.sync_mode")
	if syncMode == "" {
		syncMode = string(SyncModeDefault)
	}

	pn := NewPeerNode(nodeAddr, nodePort, common.HexToAddress(opts.Address), SyncMode(syncMode))

	n := &Node{
		metadata: pn,
		httpHandler: httpServer{
			router: mux.NewRouter(),
			logger: opts.Logger,
			tracer: opts.Tracer,
		},
		config: opts,
		bc:     nil,
		logger: opts.Logger,
		knownPeers: knownPeers{
			peers: map[string]PeerNode{
				pn.Account.String(): pn,
			},
			lock: new(sync.RWMutex),
		},
		pendingTXs:    make(map[common.Hash]types.SignedTx),
		newSyncTXs:    make(chan types.SignedTx),
		newSyncBlocks: make(chan types.Block),
		//Lock:       new(sync.RWMutex),
	}

	return n, nil
}

func (n *Node) Init(ctx context.Context) error {
	sigutil.ListenExit(func(signal os.Signal) {
		n.logger.Warnf("Signal [%s] received. Graceful shutdown initialized.", signal)
		time.AfterFunc(15*time.Second, func() {
			n.logger.Fatal("Failed to gracefully shutdown after 15 sec. Force exit")
		})
		n.Shutdown()
	})

	tracer, err := traceutil.NewTracerFromViperConfig()
	if err != nil {
		if err != traceutil.ErrCollectorUrlNotSpecified {
			return err
		}
	} else {
		n.tracer = tracer
		n.httpHandler.tracer = tracer
	}

	db, err := badgerdb.OpenDB(viper.GetString("data_dir"), badger.DefaultOptions(n.config.NodeFilePath))
	if err != nil {
		n.logger.Errorf("Unable to open db file: %s", err)
		return err
	}
	n.db = db

	chain, err := core.NewBlockchain(n.config)
	if err != nil {
		n.logger.Errorf("Unable to continue blockchain: %s", err)
		return err
	} else {
		n.bc = chain
	}

	if err := chain.LoadChainState(ctx); err != nil {
		n.logger.Errorf("Unable to continue blockchain: %s", err)
		return err
	}

	n.wm, err = wallets.NewManager(n.config)
	if err != nil {
		n.logger.Errorf("Unable to init wallets manager: %s", err)
		return err
	}

	if err := n.setupNodeAccount(); err != nil {
		n.logger.Errorf("Unable to setup node account: %s", err)
		return err
	}
	n.logger.Debugf("Node account: %s", n.account.Address())

	return nil
}

func (n *Node) Run(ctx context.Context) error {
	nodeAddress := fmt.Sprintf("%s:%d", n.metadata.Ip, n.metadata.Port)
	n.logger.Infow("Starting node...",
		"addr", nodeAddress, "is_root", n.metadata.Root)
	go func() {
		n.logger.Debugw("Listening gRPC", "addr", nodeAddress)
		grpcSrv, err := n.PrepareGrpcServer()
		if err != nil {
			n.logger.Errorf("Unable to prepare gRPC server: %s", err)
			n.Shutdown()
		}
		n.grpcServer = grpcSrv

		if err := n.RunGrpcServer(nodeAddress); err != nil {
			n.logger.Errorf("Unable to start gRPC server: %s", err)
			n.Shutdown()
		}
	}()

	go n.race(ctx)
	go n.sync(ctx)

	httpApiAddress := fmt.Sprintf("%s:%s",
		viper.GetString("http.addr"),
		viper.GetString("http.port"),
	)
	n.logger.Infow("Listening HTTP", "addr", httpApiAddress)
	return n.serveHttp()
}

func (n *Node) Shutdown() {
	defer close(n.newSyncTXs)
	defer close(n.newSyncBlocks)

	var wg sync.WaitGroup

	if n.db != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := n.db.Close(); err != nil {
				n.logger.Errorf("Unable to close node db: %s", err)
			}
		}()
	}

	if n.bc != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n.bc.Shutdown()
		}()
	}

	if n.wm != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n.wm.Shutdown()
		}()
	}

	if n.grpcServer != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n.grpcServer.GracefulStop()
		}()
	}

	if n.tracer != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n.tracer.Close()
		}()
	}

	wg.Wait()

	os.Exit(0)
}

func (n *Node) IsKnownPeer(peer PeerNode) bool {
	n.knownPeers.lock.RLock()
	defer n.knownPeers.lock.RUnlock()
	_, ok := n.knownPeers.peers[peer.Account.Hex()]
	return ok
}

func (n *Node) race(ctx context.Context) {
	var miningCtx context.Context
	var stopCurrentRace context.CancelFunc

	ticker := time.NewTicker(time.Second * 10)

	for {
		select {
		case <-ticker.C:
			go func() {
				if len(n.pendingTXs) > 0 && !n.inGenRace {
					n.logger.Debugw("There is transactions available", "txs", len(n.pendingTXs))
					n.inGenRace = true

					miningCtx, stopCurrentRace = context.WithCancel(ctx)
					if err := n.generateBlock(miningCtx); err != nil {
						n.logger.Errorf("Failed to generate new block: %s", err)
					}

					n.inGenRace = false
				}
			}()

		case block, _ := <-n.newSyncBlocks:
			n.logger.Debugw("Proposed block appeared", "is_racing", n.inGenRace)
			if n.inGenRace {
				n.logger.Warnf("Another peer has won the game '%s'! :(", block.BlockHeader.Hash.Hex())
				n.removeAppliedPendingTXs(&block)
				stopCurrentRace()
			}

		case <-ctx.Done():
			ticker.Stop()
			n.logger.Debug("Party context cancelled")
			break
		}
	}
}
