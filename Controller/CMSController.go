package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	service "system-y-format-generator-faysal-cms/Service"

	"github.com/gorilla/mux"
)

type RequestParams struct {
	TransactionCode         int
	ExistingTransactionCode int
	ReferenceId             int
	AlertActivityTypeKey    string
	FromApproved            bool
}

func ParseRequestParams(r *http.Request) (*RequestParams, error) {
	params := mux.Vars(r)

	required := []string{"transactioncode", "existingTransactionCode", "referenceId", "alertActivityTypeKey", "approvedFlag"}
	for _, key := range required {
		if val, ok := params[key]; !ok || val == "" {
			log.Printf("Missing or empty parameter: %s", key)
			return nil, fmt.Errorf("missing or empty parameter: %s", key)
		}
	}

	transactionCode, err := strconv.Atoi(params["transactioncode"])
	if err != nil {
		log.Printf("Invalid transactioncode: %v", err)
		return nil, fmt.Errorf("invalid transactioncode")
	}

	referenceId, err := strconv.Atoi(params["referenceId"])
	if err != nil {
		log.Printf("Invalid referenceId: %v", err)
		return nil, fmt.Errorf("invalid referenceId")
	}

	fromApproved, err := strconv.ParseBool(params["approvedFlag"])
	if err != nil {
		log.Printf("Invalid approvedFlag: %v", err)
		return nil, fmt.Errorf("invalid approvedFlag")
	}

	return &RequestParams{
		TransactionCode:      transactionCode,
		ReferenceId:          referenceId,
		AlertActivityTypeKey: params["alertActivityTypeKey"],
		FromApproved:         fromApproved,
	}, nil
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("JSON encoding error: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetFacilityDetails(db *sql.DB, dbApproved *sql.DB, reportTable string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := ParseRequestParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Fetching Facility Details for transaction %d", req.TransactionCode)
		data := service.GetFacilityDetails(db, dbApproved, req.TransactionCode, req.ReferenceId, req.AlertActivityTypeKey, req.FromApproved, reportTable)
		writeJSON(w, data)
	}
}

func GetApproversCommentsDetails(db *sql.DB, dbApproved *sql.DB, reportTable string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := ParseRequestParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Fetching Approver Comments for transaction %d", req.TransactionCode)
		data := service.GetApproversCommentsDetails(db, dbApproved, req.TransactionCode, req.ReferenceId, req.AlertActivityTypeKey, req.FromApproved, reportTable)
		writeJSON(w, data)
	}
}

func GetTermsAndConditionsDetails(db *sql.DB, dbApproved *sql.DB, reportTable string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := ParseRequestParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Fetching Terms and Conditions for transaction %d", req.TransactionCode)
		data := service.GetTermsAndConditionsDetails(db, dbApproved, req.TransactionCode, req.AlertActivityTypeKey, req.ReferenceId, req.FromApproved, reportTable)
		writeJSON(w, data)
	}
}

func GetDeferralDetails(db *sql.DB, dbApproved *sql.DB, reportTable string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := ParseRequestParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Fetching Deferral Details for transaction %d", req.TransactionCode)
		data := service.GetDeferralDetails(db, dbApproved, req.TransactionCode, req.ReferenceId, req.AlertActivityTypeKey, req.FromApproved, reportTable)
		writeJSON(w, data)
	}
}

func GetWaiversDetails(db *sql.DB, dbApproved *sql.DB, reportTable string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := ParseRequestParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Fetching Waivers Details for transaction %d", req.TransactionCode)
		data := service.GetWaiversDetails(db, dbApproved, req.TransactionCode, req.ReferenceId, req.AlertActivityTypeKey, req.FromApproved, reportTable)
		writeJSON(w, data)
	}
}

func GetOpenExceptionDetails(db *sql.DB, dbApproved *sql.DB, reportTable string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := ParseRequestParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Fetching Open Exception Details for transaction %d", req.TransactionCode)
		data := service.GetOpenExceptionDetails(db, dbApproved, req.TransactionCode, req.ReferenceId, req.AlertActivityTypeKey, req.FromApproved, reportTable)
		writeJSON(w, data)
	}
}

func GetNewRMGObservationsDetails(db *sql.DB, dbApproved *sql.DB, reportTable string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := ParseRequestParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Fetching RMG Observations for transaction %d", req.TransactionCode)
		data := service.GetNewRMGObservationsDetails(db, dbApproved, req.TransactionCode, req.ReferenceId, req.AlertActivityTypeKey, req.FromApproved, reportTable)
		writeJSON(w, data)
	}
}
func GetTotalAssetsDetails(db *sql.DB, dbApproved *sql.DB, reportTable string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := ParseRequestParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Fetching Total Assets for transaction %d", req.TransactionCode)
		data := service.GetTotalAssetsDetails(db, dbApproved, req.TransactionCode, req.ReferenceId, req.AlertActivityTypeKey, req.FromApproved, reportTable)
		writeJSON(w, data)
	}
}
func GetMonitoringItemDetails(db *sql.DB, dbApproved *sql.DB, reportTable string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := ParseRequestParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Fetching Monitoring Items for transaction %d", req.TransactionCode)
		data := service.GetMonitoringItemDetails(db, dbApproved, req.TransactionCode, req.ReferenceId, req.AlertActivityTypeKey, req.FromApproved, reportTable)
		writeJSON(w, data)
	}
}
