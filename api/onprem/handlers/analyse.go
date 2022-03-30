package handlers

import (
	"encoding/json"
	"net/http"
)

func (l Logic) AnalyseStockHandler(w http.ResponseWriter, r *http.Request) {

}

type analyseStockRequest struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
}

func (l Logic) parseRequest(r *http.Request, payload interface{}) error {
	return json.NewDecoder(r.Body).Decode(payload)
}

func (l Logic) writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err)
}

func (l Logic) writeSuccessResponse(w http.ResponseWriter, payload interface{}) {
	w.WriteHeader(http.StatusOK)
	dc := json.NewEncoder(w)
	dc.SetIndent("", "\t")
	dc.Encode(payload)
}
