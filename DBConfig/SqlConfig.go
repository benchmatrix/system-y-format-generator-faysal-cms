package dbconfig

import (
	"database/sql"
	"os"
	models "system-y-format-generator-faysal-cms/Models"

	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/spf13/viper"
)

func GetDB(config models.Configuration) (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		config.SQLDB_Server, config.SQLDB_User, config.SQLDB_Password, config.SQLDB_Port, config.SQLDB_Name)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		return nil, fmt.Errorf("open connection failed: %w", err)
	}

	// Test the connection to ensure it's alive
	if err = conn.Ping(); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	fmt.Println("Connected to Sql Server")
	return conn, nil
}

func GetDB2(config models.Configuration) (*sql.DB, error) {
	// Connect to database

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		config.SQLDB_Server, config.SQLDB_User, config.SQLDB_Password, config.SQLDB_Port, config.SQLDB_Name1)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		return nil, fmt.Errorf("open connection failed: %w", err)
	}

	// Test the connection to ensure it's alive
	if err = conn.Ping(); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	fmt.Println("Connected to Sql Server ")

	return conn, nil
}

// func GetDB(config models.Configuration) (*sql.DB, error) {

// 	fmt.Printf("SQL DB_SERVER : " + config.SQLDB_Server)
// 	fmt.Printf("SQL DB_Password : " + config.SQLDB_Password)
// 	fmt.Printf("SQL DB_Port :" + config.SQLDB_Port)
// 	fmt.Printf("Server Address : " + config.ServerAddress)

// 	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
// 		config.SQLDB_Server, config.SQLDB_User, config.SQLDB_Password, config.SQLDB_Port, config.SQLDB_Name)
// 	db, err := sql.Open("mssql", connString)
// 	if err != nil {
// 		log.Fatal("Open connection failed:", err.Error())
// 		return nil, err
// 	}
// 	// Set appropriate connection pool settings
// 	db.SetMaxOpenConns(config.MAX_SQLDBCONN)         // Maximum open connections
// 	db.SetMaxIdleConns(config.MAX_IDLECONN)          // Maximum idle connections
// 	db.SetConnMaxLifetime(config.SQLDBCONN_LIFETIME) // Connection lifetime
// 	fmt.Printf("Connected!\n")

// 	errPing := db.Ping()
// 	if errPing != nil {
// 		db.Close()
// 		return nil, errPing
// 	}

// 	return db, nil
// }

func LoadConfig(path string) (config models.Configuration, err error) {
	viper.AddConfigPath(path)
	//change name to app_prod for production
	os.Setenv("Active_Profile", "app_faysalLocal")
	viper.SetConfigName(os.Getenv("Active_Profile"))
	fmt.Println(os.Getenv("Active_Profile") + " Selected")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
