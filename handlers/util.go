package handlers

import (
	"encoding/json"
	"net/http"
)

func (l Handlers) writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err)
}

func (l Handlers) writeSuccessResponse(w http.ResponseWriter, payload interface{}) {
	w.WriteHeader(http.StatusOK)
	dc := json.NewEncoder(w)
	dc.SetIndent("", "\t")
	dc.Encode(payload)
}

func (l Handlers) parseRequest(r *http.Request, payload interface{}) error {
	return json.NewDecoder(r.Body).Decode(payload)
}
