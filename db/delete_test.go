package db

import (
	"testing"
)

// TestDelete
// inserts some fake data into the store, and then for every item in the store, tests if deleteing it is possible
// TODO: test with a code that does not exist
func TestDelete(t *testing.T) {
	items := []ProduceItem{
		{
			ProduceCode: "delete-1",
			Name:        "",
			UnitPrice:   0.0,
		}, {
			ProduceCode: "delete-2",
			Name:        "",
			UnitPrice:   0.0,
		}, {
			ProduceCode: "delete-3",
			Name:        "",
			UnitPrice:   0.0,
		}, {
			ProduceCode: "delete-4",
			Name:        "",
			UnitPrice:   0.0,
		},
	}

	unchangedStore := append(store, items...)
	store = append(store, items...)

	for i := range store {
		testDelete(t, i)
		store = unchangedStore
	}
}

func testDelete(t *testing.T, index int) {
	status, err := Delete(store[index].ProduceCode)

	if err != nil {
		t.Fatalf("Expected delete of item to pass without error, got error: %s\n", err)
	}

	if status != Ok {
		t.Fatalf("Expcted status of OK, but got: %s\n", getErrorStatus(status))
	}
}
