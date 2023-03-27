package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
	"github.com/rockwell-uk/go-logger/logger"
)

var (
	driverConns = map[string]*sqlite3.SQLiteConn{}
)

type SQLite struct {
	config Config
	dbs    map[string]*sqlx.DB
}

func (e *SQLite) Connect() error {
	logger.Log(
		logger.LVL_APP,
		"Connecting databases\n",
	)

	sql.Register(
		"sqlite3_with_spatialite",
		&sqlite3.SQLiteDriver{
			Extensions: []string{"mod_spatialite"},
		},
	)

	e.dbs = make(map[string]*sqlx.DB)

	for _, layerType := range LayerTypes {
		dbFilePath := fmt.Sprintf("%v/%v.db", e.config.StorageFolder, layerType)
		if _, err := os.Stat(dbFilePath); errors.Is(err, os.ErrNotExist) {
			return err
		}

		db, err := sqlx.Connect("sqlite3_with_spatialite", dbFilePath)
		if err != nil {
			panic(err)
		}
		e.dbs[layerType] = db
	}

	logger.Log(
		logger.LVL_DEBUG,
		"Connected\n",
	)

	return nil
}

func (e SQLite) Stop() {
	for _, layerType := range LayerTypes {
		// Delete the driver conn
		delete(driverConns, layerType)

		// Close the database
		e.dbs[layerType].Close()
	}
}

func (e SQLite) GetDB(layerType string) *sqlx.DB {
	return e.dbs[layerType]
}

func (e SQLite) GetTableName(layerName, square string) string {
	return square
}

func (e SQLite) GetGeomFn() string {
	return "AsWKT"
}
