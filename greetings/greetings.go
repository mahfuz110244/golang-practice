// Hello returns a greeting for the named person.
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
