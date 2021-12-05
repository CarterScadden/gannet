package produce

import (
	"fmt"
	"math"
	"regexp"
)

var (
	produceCodeExpression = regexp.MustCompile(`[A-Z\d]{4}-[A-Z\d]{4}-[A-Z\d]{4}-[A-Z\d]{4}`)
	// this allows for space on the end of the string, no mention of this needing to be handled,
	// but to handle this i would just add a trimEnd to the end of the string
	// The reason this handles alphanumeric + spaces is because the given example data has
	// multiple names with spaces in them: ["Green Pepper", "Gala Apple"]
	// ASSUMPTION
	// inside of New
	produceNameExpression = regexp.MustCompile(`[\w\d]+[\w\d\s]*`)
	// for sanity check on floats
	priceExpression = regexp.MustCompile(`\d+\.\d{2}`)
)

// ProduceItem
// the struct for what a produce item looks like
// *IMPORTANT*
// create produce items with the db.produce.New(...)
// for validation of data
type ProduceItem struct {
	Name        string  `json:"name"`
	ProduceCode string  `json:"produceCode"`
	UnitPrice   float32 `json:"unitPrice"`
}

func New(name, code string, price float32) (*ProduceItem, error) {
	if err := ValidateName(name); err != nil {
		return nil, err
	}

	if err := ValidateProduceCode(code); err != nil {
		return nil, err
	}

	// ASSUMPTION
	// numbers round to nearest
	p := float32(math.Round(100*float64(price)) / 100)

	if err := ValidateUnitPrice(p); err != nil {
		return nil, err
	}

	return &ProduceItem{
		Name:        name,
		ProduceCode: code,
		UnitPrice:   p,
	}, nil
}

func ValidateName(name string) error {
	if !produceNameExpression.Match([]byte(name)) {
		return fmt.Errorf("Name: `%s` is not alphanumeric", name)
	}

	return nil
}

func ValidateProduceCode(code string) error {
	if !produceCodeExpression.Match([]byte(code)) {
		return fmt.Errorf("ProduceCode: `%s` is not valid", code)
	}

	return nil
}

func ValidateUnitPrice(price float32) error {
	// sanity check on rounded float
	// float should have 2 numbers on end and a decimal
	if !priceExpression.Match([]byte(fmt.Sprintf("%f", price))) {
		return fmt.Errorf("UnitPrice: `%f` does not pass %s", price, priceExpression.String())
	}

	return nil
}
