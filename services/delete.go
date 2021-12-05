package services

import "fmt"

// Delete
// function to remove a ProduceItem from the services.store by the given produceCode
func Delete(produceCode string) (int, error) {
	for i, p := range store {
		if p.ProduceCode == produceCode {
			store = append(store[:i], store[i+1:]...)
			return Ok, nil
		}
	}

	return NotFound, fmt.Errorf("ProduceCode: \"%s\" not found in db", produceCode)
}
