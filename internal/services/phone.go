package services

// Testo@1324#Testo@1324#
// BVPSBYBSD15G64F6XBD9N33M

import (
	"fmt"
	"os"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendMessagePhone() {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody("Hello from Golang!")

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}

// export TWILIO_ACCOUNT_SID=xxxxxxxxx
// export TWILIO_AUTH_TOKEN=xxxxxxxxx
// export TWILIO_PHONE_NUMBER=xxxxxxxxx
// export TO_PHONE_NUMBER=xxxxxxxxx
