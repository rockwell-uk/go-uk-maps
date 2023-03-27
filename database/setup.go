package database

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

//nolint:ireturn,nolintlint
func Start(config Config) StorageEngine {
	var e StorageEngine

	switch config.Engine {
	case "mysql":
		e = &MySQL{
			config: config,
		}

		// Connect to the database
		err := e.Connect()
		if err != nil {
			panic(err)
		}

	case "sqlite":
		e = &SQLite{
			config: config,
		}

		// Connect to the database
		err := e.Connect()
		if err != nil {
			panic(err)
		}

	case "pgsql":
		e = &PgSQL{
			config: config,
		}

		// Connect to the database
		err := e.Connect()
		if err != nil {
			panic(err)
		}

	default:
		panic("database engine type not set")
	}

	return e
}
