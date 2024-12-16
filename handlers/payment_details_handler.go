package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"processout-coding-challenge-tash-canter/services"

	"github.com/gorilla/mux"
)

func GetPaymentDetailsHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    paymentID := vars["id"]

    paymentDetails, err := services.GetPaymentDetails(paymentID)
    if err != nil {
		fmt.Println(err)
        http.Error(w, fmt.Sprintf("Failed to retrieve payment details for %s", paymentID), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(paymentDetails)
}
