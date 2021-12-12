package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Main Method
func main() {
	queryParam := make(map[string]interface{})
	queryParam["name"] = "estimate"
	queryParam["description"] = "estimate"
	queryParam["active"] = "true"
	queryParam["order_number"] = "10"

	queryString := ""
	for key, value := range queryParam {
		fmt.Println(key, value)
	}
	fmt.Println(queryString)

	if queryParam["description"] != "" {
		query := fmt.Sprintf(" AND description='%v'", queryParam["description"])
		fmt.Println(query)
	}
	paramString := "AND name='estimate' AND description='Estimate' AND active=false AND order_number=0 AND name=estimate AND description=Estimate AND active=false AND order_number=0"
	getStatus := `SELECT id, name, description, active, order_number
	FROM status
	WHERE deleted_at IS NULL
	ORDER BY order_number, created_at, updated_at OFFSET $1 LIMIT $2`
	if paramString != "" {
		getStatus = fmt.Sprintf("SELECT id, name, description, active, order_number FROM status WHERE deleted_at IS NULL%s ORDER BY order_number, created_at, updated_at OFFSET $1 LIMIT $2", paramString)
	}
	fmt.Println(getStatus)
	mJson, err := json.Marshal(queryParam)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonStr := string(mJson)
	fmt.Println("The JSON data is:")
	fmt.Println(jsonStr)

	invoiceNo := "IN1000024"
	otp := 1000
	otpDuration := 2
	msg := fmt.Sprintf("Your Account payment verification code for %s is %d. This code will expire in %d minutes. Please do not share your OTP with others.", invoiceNo, otp, otpDuration)
	fmt.Println(msg)
}
