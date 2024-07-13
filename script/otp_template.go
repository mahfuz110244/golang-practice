package main

import (
	"fmt"
	"strings"
)

func main() {
	// email body template with a fixed variable OTP value
	template := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>OTP from Bohubrihi Team</title>
</head>
<body>
    <h1>OTP: %s</h1>
    <p>This is your one-time password (OTP) for logging in to your account.</p>
</body>
</html>
`
	// OTP value to be replaced in the template
	otp := "123456"

	// replace the %s placeholder in the template with the OTP value
	body := fmt.Sprintf(template, otp)

	// print the final email body
	fmt.Println(strings.TrimSpace(body))
}
