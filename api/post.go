package api

import (
	"net/http"

	"go-uk-maps/api/types"

	"github.com/rockwell-uk/go-geos-draw/geom"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	var tileRequest types.TileRequest

	_, err := validateRequest(&tileRequest, r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	tileGeom, _ := geom.BoundsGeom(
		0.0,
		tileRequest.TileWidth,
		0.0,
		tileRequest.TileHeight,
	)
	tileRequest.TileGeom = tileGeom

	tile, err := makeTile(db, tileRequest)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeTile(tile, tileRequest, w, r)
}
