package handlers

import (
	"net/http"
	"time"
)

var klines []int

func (l Logic) AddStatusHandler(w http.ResponseWriter, r *http.Request) {
	go func() {
		for {
			time.Sleep(time.Second)
			klines = append(klines, len(klines))
		}
	}()

	l.writeSuccessResponse(w, map[string]string{
		"success": "simon",
	})
}

func (l Logic) CurrentStatusHandler(w http.ResponseWriter, r *http.Request) {
	l.writeSuccessResponse(w, klines)
}
