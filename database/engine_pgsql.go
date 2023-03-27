//nolint:dupl
package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rockwell-uk/go-logger/logger"
)

type PgSQL struct {
	config Config
	DB     *sqlx.DB
}

func (e *PgSQL) Connect() error {
	logger.Log(
		logger.LVL_APP,
		"Connecting to database\n",
	)

	// Configure DSN
	dsn := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v connect_timeout=%v sslmode=disable",
		e.config.User,
		e.config.Pass,
		e.config.Host,
		e.config.Port,
		e.config.Schema,
		e.config.Timeout,
	)

	// Connect
	db, err := sqlx.Connect("postgres", dsn)
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

func (e PgSQL) Stop() {
	e.DB.Close()
}

func (e PgSQL) GetDB(layerType string) *sqlx.DB {
	return e.DB
}

func (e PgSQL) GetTableName(layerName, square string) string {
	return fmt.Sprintf("%s.%s", layerName, square)
}

func (e PgSQL) GetGeomFn() string {
	return "ST_AsText"
}
