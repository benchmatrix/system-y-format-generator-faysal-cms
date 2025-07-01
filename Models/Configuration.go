package models

import "time"

type Configuration struct {
	SQLDB_Server       string        `mapstructure:"SQLDB_SERVER"`
	SQLDB_Name         string        `mapstructure:"SQLDB_NAME"`
	ServerAddress      string        `mapstructure:"SERVER_ADDRESS"`
	SQLDB_User         string        `mapstructure:"SQLDB_USER"`
	SQLDB_Port         string        `mapstructure:"SQLDB_PORT"`
	SQLDB_Password     string        `mapstructure:"SQLDB_PASSWORD"`
	SQLDB_Name1        string        `mapstructure:"SQLDB_NAME1"`
	MAX_SQLDBCONN      int           `mapstructure:"MAX_SQLDBCONN"`
	MAX_IDLECONN       int           `mapstructure:"MAX_IDLECONN"`
	SQLDBCONN_LIFETIME time.Duration `mapstructure:"SQLDBCONN_LIFETIME"`
	Logger_Level       string        `mapstructure:"Logger_DebugLevel"`
	Logger_File        bool          `mapstructure:"Logger_FileLog"`
	ReportTable        string        `mapstructure:"REPORT_TABLE"`
}
