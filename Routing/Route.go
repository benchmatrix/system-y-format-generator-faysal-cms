package routing

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	controller "system-y-format-generator-faysal-cms/Controller"
	dbconfig "system-y-format-generator-faysal-cms/DBConfig"
	"system-y-format-generator-faysal-cms/authMiddleWare"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
)

// ✅ Global DB pointers for external access (e.g., DbConnector)
var Db *sql.DB
var DbApproved *sql.DB

type HandlerContext struct {
	DB          *sql.DB
	DBApproved  *sql.DB
	ReportTable string
}

func InitRouter(ctx HandlerContext) http.Handler {
	fmt.Println("Request Handler Initialized")

	myRouter := mux.NewRouter().StrictSlash(true)
	api := myRouter.PathPrefix("/").Subrouter()
	api.Use(authMiddleWare.AuthMiddleWare(ctx.DB))

	const basePath = "/config/cms/formatgeneratorfaysal"

	api.HandleFunc(basePath+"/getHealth", controller.Health()).Methods("GET")
	api.HandleFunc(basePath+"/facilitydetails/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetFacilityDetails(ctx.DB, ctx.DBApproved, ctx.ReportTable)).Methods("GET")
	api.HandleFunc(basePath+"/approversComments/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetApproversCommentsDetails(ctx.DB, ctx.DBApproved, ctx.ReportTable)).Methods("GET")
	api.HandleFunc(basePath+"/termsandconditiondetails/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetTermsAndConditionsDetails(ctx.DB, ctx.DBApproved, ctx.ReportTable)).Methods("GET")
	api.HandleFunc(basePath+"/newrmgobservationsdetails/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetNewRMGObservationsDetails(ctx.DB, ctx.DBApproved, ctx.ReportTable)).Methods("GET")
	api.HandleFunc(basePath+"/policydeferraldetails/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetDeferralDetails(ctx.DB, ctx.DBApproved, ctx.ReportTable)).Methods("GET")
	api.HandleFunc(basePath+"/policywaiversdetails/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetWaiversDetails(ctx.DB, ctx.DBApproved, ctx.ReportTable)).Methods("GET")
	api.HandleFunc(basePath+"/policyopenexceptionsdetails/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetOpenExceptionDetails(ctx.DB, ctx.DBApproved, ctx.ReportTable)).Methods("GET")
	api.HandleFunc(basePath+"/totalassetdetails/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetTotalAssetsDetails(ctx.DB, ctx.DBApproved, ctx.ReportTable)).Methods("GET")
	api.HandleFunc(basePath+"/monitoringitems/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetMonitoringItemDetails(ctx.DB, ctx.DBApproved, ctx.ReportTable)).Methods("GET")

	myRouter.Handle("/api/customergetservice/prometheus", promhttp.Handler())

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Authorization", "Role", "username"},
		AllowCredentials: true,
	})

	return c.Handler(myRouter)
}

func StartServer() {
	config, err := dbconfig.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	db, err := dbconfig.GetDB(config)
	if err != nil {
		log.Fatal("Primary DB connection failed:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Primary DB ping failed:", err)
	}

	dbApproved, err2 := dbconfig.GetDB2(config)
	if err2 != nil {
		log.Fatal("Approved DB connection failed:", err2)
	}
	if err2 = dbApproved.Ping(); err2 != nil {
		log.Fatal("Approved DB ping failed:", err2)
	}

	// ✅ Assign to global variables for external access
	Db = db
	DbApproved = dbApproved

	handler := InitRouter(HandlerContext{
		DB:          Db,
		DBApproved:  DbApproved,
		ReportTable: config.ReportTable,
	})

	log.Println("Starting server on", config.ServerAddress)
	log.Fatal(http.ListenAndServe(config.ServerAddress, handler))
}
