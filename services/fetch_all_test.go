package services

import (
	"gannet/services/produce"
	"testing"
)

func TestFetchAll(t *testing.T) {
	c := make(chan []produce.ProduceItem)
	go FetchAll(c)
	produceItems := <-c

	for _, a := range produceItems {
		found := false
		for _, b := range store {
			if a.Name == b.Name &&
				a.UnitPrice == b.UnitPrice &&
				a.ProduceCode == b.ProduceCode {
				found = true
				break
			}
		}

		if !found {
			t.Fatalf("FetchAll did not fetch all, failed to find the pair of %v\n", a)
		}
	}
}
