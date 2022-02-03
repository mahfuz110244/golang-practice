package sum

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestSum(t *testing.T) {
	a := 10
	b := 20
	want := 30
	sum, err := CalculateSum(a, b)
	if want != sum || err != nil {
		t.Fatalf(`CalculateSum(a, b) = %q, %v, want match for %#q, nil`, sum, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestSumZero(t *testing.T) {
	want := 0
	sum, err := CalculateSum(0, 0)
	if sum != want || err == nil {
		t.Fatalf(`CalculateSum(a, b) = %v, %v, want match for %v, nil`, sum, err, want)
	}
}
