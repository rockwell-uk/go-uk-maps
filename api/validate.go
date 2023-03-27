package api

import (
	"encoding/json"
	"io"
	"net/http"

	"go-uk-maps/api/types"
)

func validateRequest(r types.ValidatedRequest, request *http.Request) ([]byte, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return []byte{}, err
	}

	err = json.Unmarshal(body, r)
	if err != nil {
		return []byte{}, err
	}

	err = r.Validate()
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
