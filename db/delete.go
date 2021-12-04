package db

import "fmt"

// Delete
// global database function which will remove a ProduceItem from the db.store by the given
// produceCode
func Delete(produceCode string) (int, error) {
	for i, p := range store {
		if p.ProduceCode == produceCode {
			// mutate store array
			store = append(store[:i], store[i+1:]...)
			return Ok, nil
		}
	}

	return NotFound, fmt.Errorf("ProduceCode: \"%s\" not found in db", produceCode)
}
