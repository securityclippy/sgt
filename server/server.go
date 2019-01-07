package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oktasecuritylabs/sgt/dyndb"
	"github.com/oktasecuritylabs/sgt/handlers/api"
	"github.com/oktasecuritylabs/sgt/handlers/auth"
	"github.com/oktasecuritylabs/sgt/handlers/distributed"
	"github.com/oktasecuritylabs/sgt/handlers/node"
	"github.com/oktasecuritylabs/sgt/internal/pkg/filecarver"
	"github.com/oktasecuritylabs/sgt/osquery_types"
	"github.com/urfave/negroni"
	"gitlab.com/clippy/ec2autocert/pkg/config"
	"github.com/mholt/certmagic"
	"github.com/xenolf/lego/providers/dns/route53"
	"strconv"
	"github.com/xenolf/lego/log"
	"github.com/securityclippy/magicstorage"
)

// Serve will create the server listen
func Serve() error {
	dynb := dyndb.NewDynamoDB()

	router := mux.NewRouter()
	serverConfig, err := osquery_types.GetServerConfig("config.json")
	if err != nil {
		return err
	}
	//node endpoint
	nodeAPI := router.PathPrefix("/node").Subrouter()
	nodeAPI.Path("/configure").Handler(node.NodeConfigureRequest(dynb, serverConfig))
	nodeAPI.Path("/enroll").Handler(node.NodeEnrollRequest(dynb, serverConfig))
	//protect with uiAuth
	//Configuration (management) endpoint
	apiRouter := mux.NewRouter().PathPrefix("/api/v1/configuration").Subrouter()

	//apiRouter.HandleFunc("/configs", api.GetNamedConfigs).Methods(http.MethodGet, http.MethodPost)
	apiRouter.Handle("/configs", api.GetNamedConfigsHandler(dynb)).Methods(http.MethodGet, http.MethodPost)
	apiRouter.Handle("/configs/{config_name}", api.ConfigurationRequestHandler(dynb))
	//apiRouter.HandleFunc("/configs/{config_name}", api.ConfigurationRequest).Methods(http.MethodPost)
	//Nodes
	//apiRouter.HandleFunc("/nodes", api.GetNodes).Methods(http.MethodGet)
	apiRouter.Handle("/nodes", api.GetNodesHandler(dynb))
	//apiRouter.HandleFunc("/nodes/{node_key}", api.ConfigureNode).Methods(http.MethodPost, http.MethodGet)
	apiRouter.Handle("/nodes/{node_key}", api.ConfigureNodeHandler(dynb))
	apiRouter.Handle("/nodes/{node_key}", api.DeleteNodeHandler(dynb)).Methods(http.MethodDelete)
	apiRouter.Handle("/nodes/{node_key}/approve", api.ApproveNode(dynb)).Methods(http.MethodPost)
	//apiRouter.HandleFunc("/nodes/approve/_bulk", api.Placeholder).Methods("POST)
	//Packs
	apiRouter.Handle("/packs", api.GetQueryPacks(dynb)).Methods(http.MethodGet)
	apiRouter.Handle("/packs/search/{search_string}", api.SearchQueryPacks(dynb)).Methods(http.MethodGet)
	apiRouter.Handle("/packs/{pack_name}", api.ConfigurePack(dynb)).Methods(http.MethodPost)
	//PackQueries
	apiRouter.Handle("/packqueries", api.GetPackQueries(dynb)).Methods(http.MethodGet)
	apiRouter.Handle("/packqueries/{query_name}", api.ConfigurePackQuery(dynb))
	apiRouter.Handle("/packqueries/search/{search_string}", api.SearchPackQueries(dynb))
	apiRouter.Handle("/distributed/add", distributed.DistributedQueryAdd(dynb))
	//Enforce uiAuth for all our api configuration endpoints
	router.PathPrefix("/api/v1/configuration").Handler(negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(auth.AnotherValidation),
		negroni.Wrap(apiRouter),
	))
	//token
	router.Handle("/api/v1/get-token", auth.GetTokenHandler(dynb))
	//Distributed endpoint
	distributedRouter := mux.NewRouter().PathPrefix("/distributed").Subrouter()
	distributedRouter.Handle("/read", distributed.DistributedQueryRead(dynb))
	distributedRouter.Handle("/write", distributed.DistributedQueryWrite(dynb))
	//auth for distributed read/write
	router.PathPrefix("/distributed").Handler(negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(auth.ValidNodeKey),
		negroni.Wrap(distributedRouter),
	))

	carveRouter := mux.NewRouter().PathPrefix("/carve").Subrouter()
	carveRouter.Handle("/start", filecarver.StartCarve(dynb))
	carveRouter.Handle("/continue", filecarver.ContinueCarve(dynb))
	router.PathPrefix("/carve").Handler(negroni.New(
		negroni.NewRecovery(),
		//negroni.HandlerFunc(auth.ValidNodeKey),
		negroni.Wrap(carveRouter),
	))

	//Enforce auth for all our api configuration endpoints
	router.PathPrefix("/api/v1/configuration").Handler(negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(auth.AnotherValidation),
		negroni.Wrap(apiRouter),
	))

	servConfig := config.ConfigFromFile("config.json")

	dnsProvider, err := route53.NewDNSProvider()
	if err != nil {
		return err
	}


	//certmagic.DNSProvider = dn

	certmagic.DNSProvider = dnsProvider
	certmagic.DefaultStorage = magicstorage.NewS3Storage(servConfig.S3BackendBucket, "us-east-1")

	//testing to see if this relaunches

	useStaging, err := strconv.ParseBool(servConfig.UseLEStaging)
	if err != nil {
		log.Fatal(err)
	}

	if useStaging {
		certmagic.CA = certmagic.LetsEncryptStagingCA
	} else {
		certmagic.CA = certmagic.LetsEncryptProductionCA
	}


	return certmagic.HTTPS([]string{servConfig.Domain}, router)
	//m := certmanager.NewManager(servConfig, router)

	//m.ServeWithAutoUpdate(servConfig.ListenAddress, router)
	//err = http.ListenAndServeTLS(":443",
		//"fullchain.pem", "privkey.pem", router)
	//"fullchain.pem", "privkey.pem", handlers.LoggingHandler(os.Stdout, router))
	//return err
}
