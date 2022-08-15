package handlers

import (
	"net/http"

	"github.com/mcuv3/mcbot/internal/logic"
)

var _ = Handler(&Handlers{})

type Handler interface {
	AnalyzeStockHandler(w http.ResponseWriter, r *http.Request)
	AddStockHandler(w http.ResponseWriter, r *http.Request)
	ListStockHandler(w http.ResponseWriter, r *http.Request)
	DeleteStockHandler(w http.ResponseWriter, r *http.Request)
	SeedHandler(w http.ResponseWriter, r *http.Request)
	AnalyzeStateHandler(w http.ResponseWriter, r *http.Request)
}

type Params struct {
	Logic logic.Layer
}

type Handlers struct {
	logic logic.Layer
}

func NewHandlers(params Params) Handlers {
	return Handlers{
		logic: params.Logic,
	}
}
