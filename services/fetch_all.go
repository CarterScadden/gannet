package services

import "gannet/services/produce"

func FetchAll(c chan []produce.ProduceItem) {
	c <- store
	close(c)
}
