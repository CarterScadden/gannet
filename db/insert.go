package db

import (
	"fmt"
	"gannet/db/produce"
)

// Insert
// global database function which will append a ProduceItem to the db.store
// fails with Conflict if any of the given produce items conflict with the store
// or fails with BadRequest if any of the produce items fail to pass validation
// else returns an Ok, with a nil error
func Insert(ps ...produce.ProduceItem) (int, error) {
	// TODO: potentially a better solution here
	for _, p := range ps {
		for _, item := range store {
			if item.ProduceCode == p.ProduceCode {
				return Conflict, fmt.Errorf("ProduceCode: \"%s\" already exists in db, denying request", p.ProduceCode)
			}
		}
	}

	store = append(store, ps...)

	return Ok, nil
}
