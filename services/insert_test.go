package services

import (
	"gannet/services/produce"
	"net/http"
	"testing"
)

// TestInsert
// tests that Insert is working properly by
// inserting a single item and testing if it is there,
// inserting multiple items and testing if they are all there
// inserting a bad item and recieving a conflict status along with the data not being
// where it is supposed to be
func TestInsert(t *testing.T) {
	items := []produce.ProduceItem{
		{
			Name:        "",
			ProduceCode: "b",
			UnitPrice:   0.0,
		},
		{
			Name:        "",
			ProduceCode: "c",
			UnitPrice:   0.0,
		},
	}

	c := make(chan int)

	go Insert(c, items...)
	status := <-c

	if status != http.StatusOK {
		t.Fatalf("Expected insert of %v to pass with 200, instead got status: %d\n", items, status)
	}

	for _, p := range items {
		found := false

		for _, item := range store {
			if item.ProduceCode == p.ProduceCode {
				found = true
				break
			}
		}

		if !found {
			t.Fatalf("expected %s to be in store: %v", p.ProduceCode, store)
		}
	}

	// Do again to check and see if we get conflict status

	go Insert(c, items...)
	status = <-c

	if status != http.StatusConflict {
		t.Fatalf("Expected insert of %v to fail with 409, instead got status: %d\n", items, status)
	}

	for _, p := range items {
		count := 0

		for _, item := range store {
			if item.ProduceCode == p.ProduceCode {
				count++
			}
		}

		if count != 1 {
			t.Fatalf("expected %s to not have conflicting dupulicate in store, found %d copies", p.ProduceCode, count)
		}
	}

}
