package node

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/prom2json"
	"github.com/rovergulf/rbn/core"
	"github.com/rovergulf/rbn/pkg/response"
	"github.com/rovergulf/rbn/pkg/version"
	"net/http"
	"strings"
	"time"
)

func (n *Node) serveHttp() error {
	r := n.httpHandler.router

	r.HandleFunc("/health", n.healthCheck).Methods(http.MethodGet)
	r.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)
	r.HandleFunc("/metrics/json", n.DiscoverMetrics).Methods(http.MethodGet)
	r.HandleFunc("/routes", n.WalkRoutes).Methods(http.MethodGet)

	r.HandleFunc(endpointStatus, n.NodeStatus).Methods(http.MethodGet)

	r.HandleFunc("/blocks/latest", n.LatestBlock).Methods(http.MethodGet)
	r.HandleFunc("/blocks/{hash}", n.FindBlock).Methods(http.MethodGet)
	r.HandleFunc("/balances", n.ListBalances).Methods(http.MethodGet)
	r.HandleFunc("/balances/{addr}", n.GetBalance).Methods(http.MethodGet)
	r.HandleFunc("/tx/add", n.TxAdd).Methods(http.MethodGet)

	return http.ListenAndServe(n.metadata.ApiAddress(), &n.httpHandler)
}

func (n *Node) httpResponse(w http.ResponseWriter, i interface{}, statusCode ...int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if len(statusCode) > 0 {
		w.WriteHeader(statusCode[0])
	} else {
		w.WriteHeader(http.StatusOK)
	}

	if err := response.WriteJSON(w, n.logger, i); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Unable to write json response: %s", err.Error())))
	}
}

func (n *Node) healthCheck(w http.ResponseWriter, r *http.Request) {
	n.httpResponse(w, map[string]interface{}{
		"http_status": http.StatusOK,
		"timestamp":   time.Now().Unix(),
		"run_date":    version.RunDate.Format(time.RFC1123),
		"node_status": "healthy",
		"is_mining":   n.isMining,
	})
}

func (n *Node) DiscoverMetrics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	metricsUrl := fmt.Sprintf("%s/metrics", n.metadata.HttpApiAddress())
	req, err := http.Get(metricsUrl)
	if err != nil {
		n.logger.Errorf("Unable to send request to prometheus metrics: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		n.httpResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mfChan := make(chan *dto.MetricFamily, 1024)

	// Missing input means we are reading from an URL.
	if err := prom2json.ParseReader(req.Body, mfChan); err != nil {
		n.logger.Errorf("error reading metrics: %s", err)
		n.httpResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var result []*prom2json.Family
	for mf := range mfChan {
		result = append(result, prom2json.NewFamily(mf))
	}

	n.httpResponse(w, result)
}

func (n *Node) WalkRoutes(w http.ResponseWriter, r *http.Request) {
	var results []map[string]interface{}

	err := n.httpHandler.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		res := make(map[string]interface{})

		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			//h.Logger.Debug("ROUTE: ", pathTemplate)
			res["route"] = pathTemplate
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			//h.Logger.Debug("Path regexp: ", pathRegexp)
			res["regexp"] = pathRegexp
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			//h.Logger.Debug("Queries templates: ", strings.Join(queriesTemplates, ","))
			res["queries_templates"] = strings.Join(queriesTemplates, ",")
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			//h.Logger.Debug("Queries regexps: ", strings.Join(queriesRegexps, ","))
			res["queries_regexps"] = strings.Join(queriesRegexps, ",")
		}
		methods, err := route.GetMethods()
		if err == nil {
			//h.Logger.Debug("Methods: ", strings.Join(methods, ","))
			res["methods"] = methods
		}

		results = append(results, res)
		return nil
	})
	if err != nil {
		n.logger.Error(err)
	}

	n.httpResponse(w, results)
}

func (n *Node) NodeStatus(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	lb, err := n.bc.GetBlock(n.bc.LastHash)
	if err != nil {
		n.httpResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lsm, vlog := n.bc.Db.Size()
	n.httpResponse(w, StatusRes{
		LastHash:   lb.Hash.Hex(),
		Number:     n.bc.ChainLength.Uint64(),
		KnownPeers: n.knownPeers,
		PendingTXs: n.pendingTXs,
		IsMining:   n.isMining,
		DbSizeLsm:  lsm,
		DbSizeVlog: vlog,
	})
}

func (n *Node) AddPeerNode(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	n.logger.Debug("http server AddPeerNode called")
	n.httpResponse(w, true, http.StatusNotImplemented)
}

func (n *Node) SyncPeers(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	n.logger.Debug("http server SyncPeers called")
	n.httpResponse(w, true, http.StatusNotImplemented)
}

func (n *Node) ListBalances(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	n.logger.Debug("http server ListBalances called")

	balances := core.Balances{
		Blockchain: n.bc,
	}

	list, err := balances.ListBalances()
	if err != nil {
		n.httpResponse(w, err, http.StatusInternalServerError)
		return
	}

	n.httpResponse(w, list)
}

func (n *Node) GetBalance(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	vars := mux.Vars(r)
	addr := vars["addr"]

	if !common.IsHexAddress(addr) {
		n.httpResponse(w, fmt.Errorf("invalid address: %s", addr), http.StatusBadRequest)
		return
	}

	balances := core.Balances{
		Blockchain: n.bc,
	}

	balance, err := balances.GetBalance(common.HexToAddress(addr))
	if err != nil {
		n.httpResponse(w, err, http.StatusBadRequest)
		return
	}

	n.httpResponse(w, balance)
}

func (n *Node) TxAdd(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	n.logger.Debug("http server TxAdd called")
	n.httpResponse(w, true, http.StatusNotImplemented)
}

func (n *Node) LatestBlock(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	b, err := n.bc.GetBlock(n.bc.LastHash)
	if err != nil {
		n.httpResponse(w, err, http.StatusInternalServerError)
		return
	}

	n.httpResponse(w, b)
}

func (n *Node) FindBlock(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	vars := mux.Vars(r)
	hash := vars["hash"]
	if len(hash) == 0 {
		n.httpResponse(w, fmt.Errorf("invalid hash: %s", hash), http.StatusBadRequest)
		return
	}

	b, err := n.bc.GetBlock(common.HexToHash(hash))
	if err != nil {
		n.httpResponse(w, err, http.StatusInternalServerError)
		return
	}

	n.httpResponse(w, b)
}
