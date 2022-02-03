// Hello returns a greeting for the named person.
package sum

import "fmt"

// Hello returns a greeting for the named person.
func CalculateSum(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("zero params not accepted")
	}
	return a + b, nil
}
