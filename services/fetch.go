package services

import (
	"gannet/services/produce"
)

// Fetch
// global database function which will search for for a produceCode by the given code
// and return that an address to that value
// if nothing was found, the channel given is closed
func Fetch(c chan *produce.ProduceItem, code string) {
	for _, item := range store {
		if item.ProduceCode == code {
			c <- &item
			return
		}
	}

	close(c)
}
