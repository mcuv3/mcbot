package handlers

import "net/http"

func (h Handlers) SeedHandler(w http.ResponseWriter, r *http.Request) {
	if err := h.logic.Seed(r.Context()); err != nil {
		h.writeError(w, err)
		return
	}
	h.writeSuccessResponse(w, map[string]string{
		"success": "simon",
	})
}
