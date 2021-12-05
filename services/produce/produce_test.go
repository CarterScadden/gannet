package produce

import "testing"

func TestValidateName(t *testing.T) {
	shouldWork := []string{
		"Apples",
		"apples",
		"apples    ",
		"orange 2",
		"9 awesome people are awesome",
		"do not fail me now",
	}

	shouldNotWork := []string{
		"  Apples ",
		" Apples",
		"! greenbeans",
		"Lettuce#",
	}

	for _, should := range shouldWork {
		if err := ValidateName(should); err != nil {
			t.Fatal(err)
		}
	}

	for _, shouldNot := range shouldNotWork {
		if err := ValidateName(shouldNot); err != nil {
			t.Fatalf("`%s` marked as valid, when it should not be", shouldNot)
		}
	}
}
