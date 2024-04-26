package controller

import (
	"encoding/json"
	"net/http"
)

func readBodyToStruct[T any](r *http.Request, out *T) (*T, error) {
	err := json.NewDecoder(r.Body).Decode(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
