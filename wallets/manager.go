package wallets

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/rovergulf/rbn/database/badgerdb"
	"github.com/rovergulf/rbn/params"
	"github.com/rovergulf/rbn/pkg/traceutil"
	"go.uber.org/zap"
)

const DbWalletFile = "wallets.db"

type Manager struct {
	db     *badger.DB
	logger *zap.SugaredLogger
	tracer traceutil.Tracer
	quit   chan struct{}
}

// NewManager returns wallets Manager instance
func NewManager(opts params.Options) (*Manager, error) {
	badgerOpts := badger.DefaultOptions(opts.WalletsFilePath)
	db, err := badgerdb.OpenDB(opts.WalletsFilePath, badgerOpts)
	if err != nil {
		opts.Logger.Errorf("Unable to open db file: %s", err)
		return nil, err
	}

	return &Manager{
		db:     db,
		logger: opts.Logger,
	}, err
}

func (m *Manager) DbSize() (int64, int64) {
	return m.db.Size()
}

func (m *Manager) Shutdown() {
	if m.tracer != nil {
		m.tracer.Close()
	}

	if m.db != nil {
		if err := m.db.Close(); err != nil {
			m.logger.Errorf("Unable to close wallets db: %s", err)
		}
	}
}
