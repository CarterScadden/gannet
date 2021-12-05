package services

import "gannet/services/produce"

var (
	store = []produce.ProduceItem{
		{
			ProduceCode: "A12T-4GH7-QPL9-3N4M",
			Name:        "Lettuce",
			UnitPrice:   3.46,
		},
		{
			ProduceCode: "E5T6-9UI3-TH15-QR88",
			Name:        "Peach",
			UnitPrice:   2.99,
		},
		{

			ProduceCode: "YRT6-72AS-K736-L4AR",
			Name:        "Green Pepper",
			UnitPrice:   0.79,
		},
		{

			ProduceCode: "TQ4C-VV6T-75ZX-1RMR",
			Name:        "Gala Apple",
			UnitPrice:   3.59,
		},
	}
)

// validate that the initial values are valid
func init() {
	for _, item := range store {
		if err := produce.ValidateName(item.Name); err != nil {
			panic(err)
		}

		if err := produce.ValidateProduceCode(item.ProduceCode); err != nil {
			panic(err)
		}

		if err := produce.ValidateUnitPrice(item.UnitPrice); err != nil {
			panic(err)
		}
	}
}
