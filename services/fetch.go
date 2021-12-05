package services

import "gannet/db/produce"

// Fetch
// global database function which will grab a ProduceItem from the db.store by the given
// produceCode
func Fetch(produceCodes ...string) []produce.ProduceItem {
	// TODO: potentially a better solution here
	items := []produce.ProduceItem{}

	for _, code := range produceCodes {
		for _, item := range store {
			if code == item.ProduceCode {
				items = append(items, item)
				break
			}
		}
	}

	return items
}