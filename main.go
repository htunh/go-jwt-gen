package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define a struct to hold the payload
type Payload struct {
	MerchantID     string   `json:"merchantID"`
	InvoiceNo      string   `json:"invoiceNo"`
	Description    string   `json:"description"`
	Amount         float64  `json:"amount"`
	CurrencyCode   string   `json:"currencyCode"`
	PaymentChannel []string `json:"paymentChannel"`
}

// Define a struct to hold the JWT payload
type JWTPayload struct {
	Payload string `json:"payload"`
}

// Function to decode the JWT token
func decodeJWT(tokenString, secretKey string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func main() {
	secretKey := "CD229682D3297390B9F66FF4020B758F4A5E625AF4992E5D75D311D6458B38E2"
	payload := Payload{
		MerchantID:     "JT04",
		InvoiceNo:      "123456789097",
		Description:    "item 1",
		Amount:         1000.00,
		CurrencyCode:   "THB",
		PaymentChannel: []string{"CC"},
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"merchantID":     payload.MerchantID,
		"invoiceNo":      payload.InvoiceNo,
		"description":    payload.Description,
		"amount":         payload.Amount,
		"currencyCode":   payload.CurrencyCode,
		"paymentChannel": payload.PaymentChannel,
		"exp":            time.Now().Add(time.Hour * 1).Unix(), // Token expiration time
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatalf("Error signing token: %v", err)
	}

	// Prepare the data to be sent
	jwtPayload := JWTPayload{Payload: tokenString}
	jsonData, err := json.Marshal(jwtPayload)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// Send the JWT via HTTP POST
	url := "https://sandbox-pgw.2c2p.com/payment/4.3/paymentToken"
	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/*+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response
	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	responseJSON, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling response JSON: %v", err)
	}

	fmt.Println("Not Decoded Response:")
	fmt.Println(string(responseJSON))

	// Extract the JWT token from the response
	encodedToken := responseBody["payload"].(string)
	decodedClaims, err := decodeJWT(encodedToken, secretKey)
	if err != nil {
		log.Fatalf("Error decoding JWT token: %v", err)
	}

	// Print the decoded claims
	decodedJSON, err := json.MarshalIndent(decodedClaims, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling decoded claims to JSON: %v", err)
	}

	fmt.Println("Decoded Response:")
	fmt.Println(string(decodedJSON))
	fmt.Println(decodedClaims["merchantID"])
}
