package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tonimnim/GO-C2B/pkg/mpesa"
)

func setupServer(daraja mpesa.DarajaAPI, stkPush mpesa.LipaNaMpesaPayload) {
	r := gin.Default()

	// Setup callback route
	callbackGroup := r.Group("/")
	mpesa.MapExpressGinCallBack(callbackGroup, "callback", func(response *mpesa.CallBackResponse, request http.Request, err error) {
		if err != nil {
			log.Printf("Callback error: %v", err)
			return
		}
		log.Printf("Payment callback received: %+v", response)
	})

	// Add STK push endpoint
	r.POST("/stkpush", func(c *gin.Context) {
		response, err := daraja.MakeSTKPushRequest(stkPush)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	r.Run(":" + port)
}
