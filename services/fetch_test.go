package services

import (
	"gannet/services/produce"
	"testing"
)

func TestFetch(t *testing.T) {
	// TODO: move this array, and the fetchTestProduceItemsEqual to some test utils file
	items := []produce.ProduceItem{
		{
			ProduceCode: "fetch-1",
			Name:        "",
			UnitPrice:   0.0,
		}, {
			ProduceCode: "fetch-2",
			Name:        "",
			UnitPrice:   0.0,
		}, {
			ProduceCode: "fetch-3",
			Name:        "",
			UnitPrice:   0.0,
		}, {
			ProduceCode: "fetch-4",
			Name:        "",
			UnitPrice:   0.0,
		},
	}

	store = append(store, items...)

	for _, item := range store {
		fetched := Fetch(item.ProduceCode)

		length := len(fetched)

		if length != 1 {
			t.Fatalf("number of fetched items for \"Fetch(item.ProduceCode)\" is not 1, got: %d\n", length)
		}

		if !fetchTestProduceItemsEqual(fetched[0], item) {
			t.Fatalf("fetched item: %v != %v\n", fetched, item)
		}
	}

	codes := []string{}

	for _, item := range store {
		codes = append(codes, item.ProduceCode)
	}

	codesToFetch := codes[1:]
	fetched := Fetch(codesToFetch...)

	for _, code := range codesToFetch {
		found := false
		for _, item := range fetched {
			if code == item.ProduceCode {
				found = true
				break
			}
		}

		if !found {
			t.Fatalf("Failed to fetch code %s\n", code)
		}
	}
}

func fetchTestProduceItemsEqual(a, b produce.ProduceItem) bool {
	return a.Name == b.Name &&
		a.UnitPrice == b.UnitPrice &&
		a.ProduceCode == b.ProduceCode
}
