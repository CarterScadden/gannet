package services

import (
	"gannet/services/produce"
	"testing"
)

// TestInsert
// test that the insert is working properly by
// inserting a single item and testing if it is there,
// inserting multiplt items and testing if they are all there
// TODO: test with a bad produce item
func TestInsert(t *testing.T) {
	a := produce.ProduceItem{
		Name:        "",
		ProduceCode: "a",
		UnitPrice:   0.0,
	}

	status, err := Insert(a)
	testInsertResult(t, status, err, a)

	variadicArguments := []produce.ProduceItem{
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

	status, err = Insert(variadicArguments...)
	testInsertResult(t, status, err, variadicArguments...)

}

func testInsertResult(t *testing.T, status int, err error, ps ...produce.ProduceItem) {
	if err != nil {
		t.Fatalf("Expected insert of 1 item to pass without error, got error: %s\n", err)
	}

	if status != Ok {
		t.Fatalf("Expcted status of OK, but got: %s\n", getErrorStatus(status))
	}

	// check that the given produce items are in the store
	for _, p := range ps {
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
}