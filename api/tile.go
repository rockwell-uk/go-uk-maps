package api

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"strconv"

	"github.com/rockwell-uk/go-geos-draw/geom"
	"github.com/rockwell-uk/go-logger/logger"

	"go-uk-maps/api/types"
	"go-uk-maps/database"
	"go-uk-maps/layerdata"
	"go-uk-maps/makeimage"
)

func tileHandler(w http.ResponseWriter, r *http.Request) {
	var tileRequest types.TileRequest
	var quality int

	_, err := validateRequest(&tileRequest, r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	tile, err := makeTile(db, tileRequest)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	buffer := new(bytes.Buffer)

	switch tileRequest.Format {
	case "png":
		// Write the image into the buffer
		if err := png.Encode(buffer, *tile); err != nil {
			msg := fmt.Sprintf("unable to encode image %v", err)
			logger.Log(
				logger.LVL_FATAL,
				msg,
			)
			writeError(w, http.StatusInternalServerError, errors.New(msg))
			return
		}

		// Write the response
		if _, err := w.Write(buffer.Bytes()); err != nil {
			msg := fmt.Sprintf("unable to write image %v", err)
			logger.Log(
				logger.LVL_FATAL,
				msg,
			)
			writeError(w, http.StatusInternalServerError, errors.New(msg))
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
		return

	default:
		if tileRequest.Quality == 0 {
			quality = 100
		} else {
			quality = tileRequest.Quality
		}

		opt := jpeg.Options{
			Quality: quality,
		}

		// Write the image into the buffer
		err = jpeg.Encode(buffer, *tile, &opt)
		if err != nil {
			msg := fmt.Sprintf("unable to encode image %v", err)
			logger.Log(
				logger.LVL_FATAL,
				msg,
			)
			writeError(w, http.StatusInternalServerError, errors.New(msg))
			return
		}

		// Write the response
		if _, err := w.Write(buffer.Bytes()); err != nil {
			msg := fmt.Sprintf("unable to write image %v", err)
			logger.Log(
				logger.LVL_FATAL,
				msg,
			)
			writeError(w, http.StatusInternalServerError, errors.New(msg))
			return
		}

		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
		return
	}
}

func writeError(w http.ResponseWriter, rcode int, err error) {
	body := []byte(err.Error())

	w.WriteHeader(rcode)
	w.Header().Set("Content-Type", "application/text")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))

	w.Write(body) //nolint:errcheck
}

func makeTile(db database.StorageEngine, r types.TileRequest) (*image.Image, error) {
	tileGeom, err := geom.BoundsGeom(
		0.0,
		r.TileWidth,
		0.0,
		r.TileHeight,
	)
	if err != nil {
		return nil, err
	}

	r.TileGeom = tileGeom

	layerOrder, layerData, err := layerdata.GetLayerData(db, r)
	if err != nil {
		return nil, err
	}

	return makeimage.DrawImage(r, layerOrder, layerData)
}
