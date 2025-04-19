package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tonimnim/GO-C2B/pkg/mpesa"
)

func main() {
	// Load .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Daraja client
	daraja := mpesa.NewDaraja(
		os.Getenv("MPESA_CONSUMER_KEY"),
		os.Getenv("MPESA_CONSUMER_SECRET"),
		mpesa.Environment(mpesa.ENVIROMENT_SANDBOX),
	)

	// Example STK Push request
	stkPush := mpesa.LipaNaMpesaPayload{
		BusinessShortCode: os.Getenv("MPESA_SHORTCODE"),
		Password:          os.Getenv("MPESA_PASSKEY"),
		TransactionType:   "CustomerPayBillOnline",
		Amount:            "1",
		PartyA:            "254759909017",
		PartyB:            os.Getenv("MPESA_SHORTCODE"),
		PhoneNumber:       "254759909017",
		CallBackURL:       os.Getenv("CALLBACK_BASE_URL") + "/callback",
		AccountReference:  "Test",
		TransactionDesc:   "Test Payment",
	}

	setupServer(daraja, stkPush)
}
