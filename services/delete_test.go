package services

import (
	"gannet/services/produce"
	"net/http"
	"testing"
)

// TestDelete
// inserts some fake data into the store, and then for every item in the store, tests if deleteing it is possible
// TODO: test with a code that does not exist
func TestDelete(t *testing.T) {
	items := []produce.ProduceItem{
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

	store = append(store, items...)

	for _, item := range items {
		testDelete(t, item.ProduceCode)
	}
}

func testDelete(t *testing.T, code string) {
	c := make(chan int)
	go Delete(c, code)
	status, ok := <-c

	if !ok {
		t.Fatal("Expected delete of item to pass ok, got !ok\n")
	}

	if status != http.StatusOK {
		t.Fatalf("Expcted status of OK, but got: %d\n", status)
	}
}
