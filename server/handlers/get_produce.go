package handlers

import (
	"encoding/json"
	"gannet/services"
	"gannet/services/produce"
	"net/http"
)

// GetProduce
// endpoint handler for getting a produce item from the store
func GetProduce(w http.ResponseWriter, r *http.Request) {
	c := make(chan []produce.ProduceItem)
	go services.FetchAll(c)
	store := <-c

	data, err := json.Marshal(store)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO: if not production, send error, else send Server Error
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}
