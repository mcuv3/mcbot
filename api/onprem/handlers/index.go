package handlers

import (
	"net/http"

	"github.com/mcuv3/mcbot/internal/storage"
)

type Handler interface {
	AnalyseStockHandler(w http.ResponseWriter, r *http.Request)
	AddStockHandler(w http.ResponseWriter, r *http.Request)
	ListStockHandler(w http.ResponseWriter, r *http.Request)
	DeleteStockHandler(w http.ResponseWriter, r *http.Request)
	CurrentStatusHandler(w http.ResponseWriter, r *http.Request)
	AddStatusHandler(w http.ResponseWriter, r *http.Request)
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
