package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProcessPaymentHandler(t *testing.T) {
	payload := map[string]interface{}{
		"card_number": "4539148803436467",
		"expiry_month": 12,
		"expiry_year": 	2025,
		"cvv":        	"123",
		"amount":     	100.50,
		"currency":   	"USD",
	}
	payloadBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/payments", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	r := setupRouter()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status OK, got %v for /payments endpoint", status)
	}

	var response map[string]interface{}
    if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
        t.Fatalf("Failed to decode response body: %v", err)
    }
	if _, exists := response["id"]; !exists {
        t.Errorf("Expected payment ID in response, but it was not found")
    }
	if _, exists := response["processed_at"]; !exists {
        t.Errorf("Expected processed at in response, but it was not found")
    }
	if response["status"] != "Success" {
		t.Errorf("Expected status to be %v but got %v", "Success", response["status"])
	}
	if response["masked_card"] != "**** **** **** 6467" {
		t.Errorf("Expected masked card to be %v but got %v", "**** **** **** 6467", response["masked_card"])
	}
	if response["amount"] != payload["amount"] {
		t.Errorf("Expected amount to be %v but got %v", payload["amount"], response["amount"])
	}
	if response["currency"] != payload["currency"] {
		t.Errorf("Expected currency to be %v but got %v", payload["currency"], response["currency"])
	}

	// Step 2: Call the Payment Details handler
	paymentID := response["id"]
	getResp, err := http.Get(fmt.Sprintf("http://localhost:8080/payments/%s", paymentID))
	if err != nil {
		t.Fatalf("Error making GET request for payment details: %v", err)
	}
	defer getResp.Body.Close()

	if getResp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got: %d for /payments/${id} endpoint", getResp.StatusCode)
	}

	var getResponse map[string]interface{}
	if err := json.NewDecoder(getResp.Body).Decode(&getResponse); err != nil {
		t.Fatalf("Failed to decode payment details response: %v", err)
	}

	if getResponse["id"] != paymentID {
		t.Fatalf("Expected payment ID %s, but got: %v", paymentID, getResponse["id"])
	}
	if _, exists := getResponse["processed_at"]; !exists {
        t.Errorf("Expected processed at in response, but it was not found")
    }
	if getResponse["status"] != "Success" {
		t.Errorf("Expected status to be %v but got %v", "Success", getResponse["status"])
	}
	if getResponse["masked_card"] != "**** **** **** 6467" {
		t.Errorf("Expected masked card to be %v but got %v", "**** **** **** 6467", getResponse["masked_card"])
	}
	if getResponse["amount"] != payload["amount"] {
		t.Errorf("Expected amount to be %v but got %v", payload["amount"], getResponse["amount"])
	}
	if getResponse["currency"] != payload["currency"] {
		t.Errorf("Expected currency to be %v but got %v", payload["currency"], getResponse["currency"])
	}
}

// Test invalid card
// Test expired card
// Test missing or invalid CVV
// Test amount and currency validation
// Split tests up for the two endpoints (this would be hard because of dynamically generated payment IDs)