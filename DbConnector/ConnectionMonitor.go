package DbConnector

import (
	"fmt"
	dbconfig "system-y-format-generator-faysal-cms/DBConfig"
	"system-y-format-generator-faysal-cms/GlobalHolder"
	models "system-y-format-generator-faysal-cms/Models"
	routing "system-y-format-generator-faysal-cms/Routing"
)

func DatabaseMonitor(config models.Configuration) {
	fmt.Println("Database Monitor Function Started..")
	for {
		select {
		case <-GlobalHolder.ReconnectSignal:
			fmt.Println("ReConnection to database signal received.")
			GlobalHolder.IsReconnecting = true
			newDB, errConnect := dbconfig.GetDB(config)
			if errConnect != nil {
				GlobalHolder.IsReconnecting = false
				GlobalHolder.ReconnectionWait <- "Db Connection Error"
			} else if newDB != nil {
				routing.Db.Close() // Close the existing connection
				routing.Db = newDB
				GlobalHolder.IsReconnecting = false
				GlobalHolder.ReconnectionSuccess <- "Reconnected"
			} else {
				GlobalHolder.IsReconnecting = false
				GlobalHolder.ReconnectionWait <- "Db Connection Error"
			}

		}

	}
}
