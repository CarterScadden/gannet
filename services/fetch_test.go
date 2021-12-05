package services

import (
	"fmt"
	"gannet/services/produce"
	"testing"
)

func TestFetch(t *testing.T) {
	c := make(chan *produce.ProduceItem)

	for _, item := range store[0:4] {
		fmt.Printf("looking for: %v\n", item.ProduceCode)
		go Fetch(c, item.ProduceCode)
		found, ok := <-c

		if found == nil {
			t.Fatalf("go Fetch(c, %s) gave an unexpected nil response\n", item.ProduceCode)
		}

		if !ok {
			t.Fatalf("go Fetch(c, %s) gave a unexpected !ok response\n", item.ProduceCode)
		}

		if !fetchTestProduceItemsEqual(*found, item) {
			t.Fatalf("fetched item: %v != %v\n", *found, item)
		}
	}

	go Fetch(c, "<id that should result in a channel close>")
	found, ok := <-c

	if ok {
		t.Fatal("Got Unexpected ok, when looking for expected NotFound\n")
	}

	if found != nil {
		t.Fatalf("Got Unexpected value, when looking for expected NotFound. Got value: `%v`\n", *found)
	}
}

func fetchTestProduceItemsEqual(a, b produce.ProduceItem) bool {
	return a.Name == b.Name &&
		a.UnitPrice == b.UnitPrice &&
		a.ProduceCode == b.ProduceCode
}
