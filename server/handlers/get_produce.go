package handlers

import (
	"encoding/json"
	"gannet/services"
	"net/http"
)

// GetProduce
// endpoint handler for getting a produce item from the store
func GetProduce(w http.ResponseWriter, r *http.Request) {
	store := services.FetchAll()

	data, err := json.Marshal(store)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO: if not production, send error, else send Server Error
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}
