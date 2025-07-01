package authMiddleWare

import (
	"database/sql"
	"net/http"
	"system-y-format-generator-faysal-cms/GlobalHolder"
)

const selectQuery = "SELECT 1"

func handleDatabaseError(w http.ResponseWriter, errMessage string) {
	http.Error(w, errMessage, http.StatusInternalServerError)
}

func DbConnectionMiddleWare(Db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := Db.Exec(selectQuery)
			if err != nil {
				var reconnectSignal <-chan string

				if GlobalHolder.IsReconnecting {
					reconnectSignal = GlobalHolder.ReconnectionSuccess
				} else {
					GlobalHolder.ReconnectSignal <- "Retry"
					reconnectSignal = GlobalHolder.ReconnectionSuccess
				}

				select {
				case <-GlobalHolder.ReconnectionWait:
					handleDatabaseError(w, "Database is currently unavailable, please try again")
					return
				case <-reconnectSignal:
					// Db Restored
					next.ServeHTTP(w, r)
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
