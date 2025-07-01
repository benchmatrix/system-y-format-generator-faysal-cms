package ErrorHandler

import (
	"context"
	"database/sql"
	"time"

	"github.com/rs/zerolog/log"
)

//func ErrorHandler(msg string, e error) {
//	if e != nil {
//		log.Error().Stack().Err(e).Msg(msg)
//		msg = msg + ":: " + e.Error()
//		writeErrorDb(conn, msg)
//		return
//	}
//	return
//
//}

func writeErrorDb(db *sql.DB, e string) {
	query := "INSERT INTO CONFIG_ERRORLINKTRANSLATION (ERRORMESSAGE) VALUES (?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Debug().Msg("Error %s when preparing SQL statement")
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, e)
	if err != nil {
		log.Debug().Msg("Error %s when inserting row into ERRORLINKTRANSLATION table")
	}
}
