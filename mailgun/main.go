package main

import (
	"context"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

// Your available domain names can be found here:
// (https://app.mailgun.com/app/domains)
var yourDomain string = "your-domain-name" // e.g. mg.yourcompany.com

// You can find the Private API Key in your Account Menu, under "Settings":
// (https://app.mailgun.com/app/account/security)
var privateAPIKey string = "your-private-key"

func main() {
	EMAIL_SENDER := "no-reply@domain.com"
	EMAIL_API_KEY := "164ab5e88576dbafe5d640b400acf81c-sample"
	DOMAIN := "domain.com"
	recipient := "mahfuzku11@gmail.com"
	otp := "123456"
	SendOtpEmail(DOMAIN, EMAIL_API_KEY, EMAIL_SENDER, recipient, otp)
}

func SendOtpEmail(domain, apiKey, sender, recipient, otp string) (string, error) {
	log.Println("Sending Email otp to recipient", recipient)
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(domain, apiKey)

	subject := "OTP Verification for Domain"

	message := mg.NewMessage(sender, subject, otp, recipient)
	body := `
		<html>
			<body>
				<p>
					Dear Sir,<br>
					<p>
						<h3>Your Domain OTP is: ` + otp + `</h3>
						<h3>Please do not share this code with anyone.</h3>
					</p>
				</p>
				<p>More details visit <a href="https://Domain.com/">https://Domain.com</a></p>
			</body>
		</html>
	`
	body = `
		<html>
			<div style="font-family: Helvetica,Arial,sans-serif;min-width:1000px;overflow:auto;line-height:2">
				<div style="margin:50px auto;width:70%;padding:20px 0">
					<div style="border-bottom:1px solid #eee">
					<a href="" style="font-size:1.4em;color: #00466a;text-decoration:none;font-weight:600">Domain</a>
					</div>
					<p style="font-size:1.1em">Hi,</p>
					<p>Thank you for choosing Domain. Use the following OTP to complete your Sign Up procedures. OTP is valid for ` + "OTP_VALIDITY" + ` minutes.
					Please do not share this code with anyone.</p>
					<h2 style="background: #00466a;margin: 0 auto;width: max-content;padding: 0 10px;color: #fff;border-radius: 4px;">` + otp + `</h2>
					<p style="font-size:0.9em;">Regards,<br />Team Domain</p>
					<hr style="border:none;border-top:1px solid #eee" />
					<div style="float:right;padding:8px 0;color:#aaa;font-size:0.8em;line-height:1;font-weight:300">
					<p>Domain</p>
					<p>House 15/B Rd No 8</p>
					<p>Dhaka 1212</p>
					</div>
				</div>
			</div>
		</html>
	`
	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ID: %s Resp: %s\n", id, resp)
	log.Println("Email sent successfully to recipient", recipient)
	return resp, err
}
