package routing

import (
	"fmt"
	"log"
	"net/http"
	controller "system-y-format-generator-faysal-cms/Controller"
	dbconfig "system-y-format-generator-faysal-cms/DBConfig"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//var Db *sql.DB

func RoutingPaths() {
	fmt.Println("Request Handler Initialized")
	//Db = DbIt

	config, errENV := dbconfig.LoadConfig(".")
	if errENV != nil {
		fmt.Println("Error in Loading Environment Variable file")
	} else {
		port := config.ServerAddress
		Db, _ := dbconfig.GetDB(config)
		DbApproved, _ := dbconfig.GetDB2(config)
		var ReportTable = config.ReportTable
		fmt.Println("ReportTable routing paths ", ReportTable)
		myRouter := mux.NewRouter().StrictSlash(true)

		//myRouter.Use(authMiddleWare.DbConnectionMiddleWare(Db))
		api := myRouter.PathPrefix("/").Subrouter()
		//api.Use(authMiddleWare.AuthMiddleWare(Db))
		api.HandleFunc("/config/cms/formatgeneratorfaysal/getHealth", controller.Health()).Methods("GET")
		api.HandleFunc("/config/cms/formatgeneratorfaysal/facilitydetails/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetFacilityDetails(Db, DbApproved, ReportTable)).Methods("GET")
		api.HandleFunc("/config/cms/formatgeneratorfaysal/approversComments/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetApproversCommentsDetails(Db, DbApproved, ReportTable)).Methods("GET")
		api.HandleFunc("/config/cms/formatgeneratorfaysal/termsandconditiondetails/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetTermsAndConditionsDetails(Db, DbApproved, ReportTable)).Methods("GET")
		api.HandleFunc("/config/cms/formatgeneratorfaysal/newrmgobservationsdetails/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetNewRMGObservationsDetails(Db, DbApproved, ReportTable)).Methods("GET")
		api.HandleFunc("/config/cms/formatgeneratorfaysal/policydeferraldetails/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetDeferralDetails(Db, DbApproved, ReportTable)).Methods("GET")
		api.HandleFunc("/config/cms/formatgeneratorfaysal/policywaiversdetails/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetWaiversDetails(Db, DbApproved, ReportTable)).Methods("GET")
		api.HandleFunc("/config/cms/formatgeneratorfaysal/policyopenexceptionsdetails/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetOpenExceptionDetails(Db, DbApproved, ReportTable)).Methods("GET")
		api.HandleFunc("/config/cms/formatgeneratorfaysal/totalassetdetails/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetTotalAssetsDetails(Db, DbApproved, ReportTable)).Methods("GET")
		api.HandleFunc("/config/cms/formatgeneratorfaysal/monitoringitems/{transactioncode}/{existingTransactionCode}/{referenceId}/{InProgress}/{approvedTransaction}/{alertActivityTypeKey}/{approvedFlag}", controller.GetMonitoringItemDetails(Db, DbApproved, ReportTable)).Methods("GET")

		handler := cors.Default().Handler(myRouter)
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedHeaders:   []string{"Authorization", "Role", "username"},
			AllowCredentials: true,
		})

		// Insert the middleware
		handler = c.Handler(handler)
		log.Fatal(http.ListenAndServe(port, handler))
	}
}
