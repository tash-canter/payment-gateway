package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"processout-coding-challenge-tash-canter/services"
)

func ProcessPaymentHandler(w http.ResponseWriter, r *http.Request) {
    var paymentRequest services.PaymentRequest
    err := json.NewDecoder(r.Body).Decode(&paymentRequest)
    if err != nil {
		fmt.Println(err)
        http.Error(w, fmt.Sprintf("Invalid payment request: %v", err), http.StatusBadRequest)
        return
    }

    result, err := services.ProcessPayment(paymentRequest)
    if err != nil {
		fmt.Println(err)
        http.Error(w, fmt.Sprintf("Payment processing failed: %v", err), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(result)
}
