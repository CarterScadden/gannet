package services

import (
	"gannet/services/produce"
	"net/http"
)

// Insert
// function to append ProduceItem's to the services.store
// pushes Conflict if any of the given produce item conflicts with the store
// else adds the items and pushes OK if everything went well
func Insert(c chan int, ps ...produce.ProduceItem) {
	for _, p := range ps {
		for _, item := range store {
			if item.ProduceCode == p.ProduceCode {
				c <- http.StatusConflict
				return
			}
		}
	}

	for _, p := range ps {
		store = append(store, p)
	}

	c <- http.StatusOK
}
