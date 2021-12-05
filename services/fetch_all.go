package services

import "gannet/services/produce"

func FetchAll() []produce.ProduceItem {
	return store
}
