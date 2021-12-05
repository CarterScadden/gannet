package services

import (
	"net/http"
)

// Delete
// function to remove a ProduceItem from the services.store by the given produceCode
func Delete(c chan int, produceCode string) {
	for i, p := range store {
		if p.ProduceCode == produceCode {
			store = append(store[:i], store[i+1:]...)
			c <- http.StatusOK
			return
		}
	}

	c <- http.StatusNotFound
}
