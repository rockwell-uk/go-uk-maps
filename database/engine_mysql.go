//nolint:dupl
package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rockwell-uk/go-logger/logger"
)

type MySQL struct {
	config Config
	DB     *sqlx.DB
}

func (e *MySQL) Connect() error {
	logger.Log(
		logger.LVL_APP,
		"Connecting to database\n",
	)

	// Configure DSN
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?timeout=%vs",
		e.config.User,
		e.config.Pass,
		e.config.Host,
		e.config.Port,
		e.config.Schema,
		e.config.Timeout,
	)

	// Connect
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	logger.Log(
		logger.LVL_DEBUG,
		"Connected\n",
	)

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Attach
	e.DB = db

	return nil
}

func (e MySQL) Stop() {
	e.DB.Close()
}

func (e MySQL) GetDB(layerType string) *sqlx.DB {
	return e.DB
}

func (e MySQL) GetTableName(layerName, square string) string {
	return fmt.Sprintf("%s.%s", layerName, square)
}

func (e MySQL) GetGeomFn() string {
	return "ST_AsWKT"
}
