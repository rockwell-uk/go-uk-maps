package layerdata

import (
	"database/sql"
)

var defaultLayers = []string{
	"administrative_boundary",

	"woodland",

	"surface_water_area",
	"surface_water_line",

	"tidal_water",
	"foreshore",
	"tidal_boundary",

	"building",

	"electricity_transmission_line",
	"glasshouse",
	"ornament",
	"railway_track",
	"railway_tunnel",

	"roundabout",
	"road_tunnel",
	"road",
	"motorway_junction",

	"railway_station",

	"named_place",

	"spot_height",
}

var layersToSimplify = []string{}

type LayerData struct {
	ID         string         `db:"id"`
	HEIGHT     sql.NullString `db:"height"`
	ROADNUMBER sql.NullString `db:"roadnumber"`
	FEATCODE   int            `db:"featcode"`
	WKT        sql.NullString `db:"wkt"`
	FONTHEIGHT sql.NullString `db:"fontheight"`
	ORIENTATIO float64        `db:"orientatio"`
	CLASSIFICA sql.NullString `db:"classifica"`
	DISTNAME   sql.NullString `db:"distname"`
	HTMLNAME   sql.NullString `db:"htmlname"`
	JUNCTNUM   sql.NullString `db:"junctnum"`
	DRAWLEVEL  sql.NullString `db:"drawlevel"`
	OVERRIDE   sql.NullString `db:"override"`
}
