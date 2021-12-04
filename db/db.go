package db

var (
	store = []ProduceItem{}
)

type ProduceItem struct {
	Name        string  `json:"name"`
	ProduceCode string  `json:"produceCode"`
	UnitPrice   float32 `json:"unitPrice"`
}

func (p ProduceItem) isValid() error {
	// TODO: validate name
	// TODO: validate produce code
	// TODO: validate price

	return nil
}
