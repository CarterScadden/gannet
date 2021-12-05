package services

import (
	"fmt"
	"net/http"
)

// Delete
// function to remove a ProduceItem from the services.store by the given produceCode
func Delete(produceCode string) (int, error) {
	for i, p := range store {
		if p.ProduceCode == produceCode {
			store = append(store[:i], store[i+1:]...)
			return http.StatusOK, nil
		}
	}

	return http.StatusNotFound, fmt.Errorf("ProduceCode: \"%s\" not found in db", produceCode)
}
