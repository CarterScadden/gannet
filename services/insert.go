package services

import (
	"fmt"
	"gannet/services/produce"
	"net/http"
)

// Insert
// function to append a ProduceItem to the services.store
// fails with Conflict if any of the given produce items conflict with the store
// or fails with BadRequest if any of the produce items fail to pass validation
// else returns an Ok, with a nil error
func Insert(ps ...produce.ProduceItem) (int, error) {
	// TODO: potentially a better solution here
	for _, p := range ps {
		for _, item := range store {
			if item.ProduceCode == p.ProduceCode {
				return http.StatusConflict, fmt.Errorf("ProduceCode: \"%s\" already exists in db, denying request", p.ProduceCode)
			}
		}
	}

	store = append(store, ps...)

	return http.StatusOK, nil
}
