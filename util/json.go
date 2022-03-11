package util

import (
	"encoding/json"
	"net/http"
)

func UnmarshalJSON(w http.ResponseWriter, r *http.Request, dto interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(&dto)
}
