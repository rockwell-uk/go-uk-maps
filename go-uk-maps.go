package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rockwell-uk/go-logger/logger"

	"go-uk-maps/api"
	"go-uk-maps/database"
	"go-uk-maps/globvar"
)

var (
	start time.Time = time.Now()

	dbengine  string = "pgsql"
	dbhost    string = "localhost"
	dbport    string = "5432"
	dbuser    string = "osdata"
	dbpass    string = "osdata"
	dbschema  string = "osdata"
	dbfolder  string = "db"
	httpport  string = "8080"
	dbtimeout int    = 10
)

func main() {
	// Define and parse flags
	var v, vv, vvv bool
	flag.BoolVar(&v, "v", false, "APP level log verbosity override")
	flag.BoolVar(&vv, "vv", false, "DEBUG level log verbosity override")
	flag.BoolVar(&vvv, "vvv", false, "INTERNAL level log verbosity override")

	// Database
	flag.StringVar(&dbengine, "dbengine", dbengine, "the database engine mysql/sqlite/pgsql")
	flag.StringVar(&dbhost, "dbhost", dbhost, "the database host")
	flag.StringVar(&dbport, "dbport", dbport, "the database port")
	flag.StringVar(&dbuser, "dbuser", dbuser, "the database username")
	flag.StringVar(&dbpass, "dbpass", dbpass, "the database password")
	flag.StringVar(&dbschema, "dbschema", dbschema, "the base database schema")
	flag.StringVar(&dbfolder, "dbfolder", dbfolder, "the storage folder for flatfile database eg. sqlite")
	flag.StringVar(&httpport, "httpport", httpport, "the port to serve HTTP on")

	flag.Parse()

	// Database
	dbConfig := database.Config{
		Engine:        dbengine,
		Host:          dbhost,
		Port:          dbport,
		User:          dbuser,
		Pass:          dbpass,
		Schema:        dbschema,
		StorageFolder: dbfolder,
		Timeout:       dbtimeout,
	}
	var db database.StorageEngine = database.Start(dbConfig)

	// Start our logger
	var vbs logger.LogLvl
	switch {
	case vvv:
		vbs = logger.LVL_INTERNAL
	case vv:
		vbs = logger.LVL_DEBUG
	case v:
		vbs = logger.LVL_APP
	}
	logger.Start(vbs)

	// Start our api
	api.HttpPort = httpport
	api.Start(db)

	// Plan for a graceful exit
	gracefulExit()
}

func gracefulExit() {
	// Setting up signal capturing
	signal.Notify(globvar.Stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-globvar.Stop

	logger.Log(
		logger.LVL_APP,
		"received sigterm",
	)

	// Stop HTTP service
	api.Stop()

	<-globvar.Http

	// Report how long the app ran for
	end := time.Now()
	diff := end.Sub(start)
	out := time.Time{}.Add(diff)

	// Log how long things took
	logger.Log(
		logger.LVL_APP,
		fmt.Sprintf("App ended %s", end.Format("15h 04m 05s")),
	)
	logger.Log(
		logger.LVL_APP,
		fmt.Sprintf("App ran for %s", out.Format("15h 04m 05s")),
	)

	// Finally stop our logger
	logger.Stop()
}
