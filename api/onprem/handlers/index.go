package handlers

import (
	"net/http"

	"github.com/MauricioAntonioMartinez/mcbot/internal/storage"
)

type Handler interface {
	AnalyseStockHandler(w http.ResponseWriter, r *http.Request)
}

type Params struct {
	Stores storage.Store
}

type Logic struct {
	Stores storage.Store
}

func NewLogic(params Params) Logic {
	return Logic{
		Stores: params.Stores,
	}
}
