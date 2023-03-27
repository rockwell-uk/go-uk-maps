package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type StorageEngine interface {
	Connect() error
	Stop()
	GetGeomFn() string
	GetDB(layerType string) *sqlx.DB
	GetTableName(layerName, square string) string
}

type Config struct {
	Engine        string
	Host          string
	Port          string
	User          string
	Pass          string
	Schema        string
	StorageFolder string
	Timeout       int
}

func (c Config) String() string {
	return fmt.Sprintf("\tEngine: %v\n\t"+
		"\tHost: %v\n\t"+
		"\tPort: %v\n\t"+
		"\tUser: %v\n\t"+
		"\tPass: %v\n\t"+
		"\tSchema: %v\n\t"+
		"\tTimeout: %v",
		c.Engine,
		c.Host,
		c.Port,
		c.User,
		c.Pass,
		c.Schema,
		c.Timeout,
	)
}

var LayerTypes = []string{
	"administrative_boundary",
	"building",
	"electricity_transmission_line",
	"foreshore",
	"functional_site",
	"glasshouse",
	"motorway_junction",
	"named_place",
	"ornament",
	"railway_station",
	"railway_track",
	"railway_tunnel",
	"road",
	"road_tunnel",
	"roundabout",
	"spot_height",
	"surface_water_area",
	"surface_water_line",
	"tidal_boundary",
	"tidal_water",
	"woodland",
}
