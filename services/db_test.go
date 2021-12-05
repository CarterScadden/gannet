package services

import "testing"

// TestStore
// test to make sure that the store is the proper length
// store is essentially tested with init
func TestStore(t *testing.T) {
	if len(store) != 4 {
		t.Fatalf("Store is not 4 long, expected initial data, got length: `%d`\n", len(store))
	}
}
