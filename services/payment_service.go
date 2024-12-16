package services

import (
	"database/sql"
	"fmt"
	"time"
)

type PaymentRequest struct {
    CardNumber 	string 	`json:"card_number"`
    ExpiryMonth int 	`json:"expiry_month"`
	ExpiryYear	int		`json:"expiry_year"`
    Amount     	float64 `json:"amount"`
    Currency   	string 	`json:"currency"`
    CVV        	string 	`json:"cvv"`
}

type PaymentResult struct {
	ID         	string  	`json:"id"`
	Status     	string  	`json:"status"`
	MaskedCard 	string  	`json:"masked_card"`
	Amount     	float64 	`json:"amount"`
	Currency   	string  	`json:"currency"`
	ProcessedAt time.Time 	`json:"processed_at"`
	ExpiryMonth int 		`json:"expiry_month"`
	ExpiryYear	int			`json:"expiry_year"`
}

func ProcessPayment(req PaymentRequest) (PaymentResult, error) {
	err := validatePaymentRequest(req)
	if err != nil {
		return PaymentResult{
			Status: "Failed",
		}, err
	}
	paymentID, status := mockBankSimulation()
	maskedCard := maskCardNumber(req.CardNumber)

	result  := PaymentResult{
		ID:         paymentID,
		Status:     status,
		MaskedCard: maskedCard,
		Amount:     req.Amount,
		Currency:   req.Currency,
		ProcessedAt: time.Now(),
		ExpiryMonth: req.ExpiryMonth,
		ExpiryYear: req.ExpiryYear,
	}
	query := `
	INSERT INTO payments (id, card_number, expiry_month, expiry_year, amount, currency, status)
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(query, paymentID, maskedCard, req.ExpiryMonth, req.ExpiryYear, req.Amount, req.Currency, status)
	if err != nil {
		return PaymentResult{
			Status: "Failed",
		}, fmt.Errorf("failed to save payment: %w", err)
	}

    return result, nil
}

func GetPaymentDetails(paymentID string) (PaymentResult, error) {
	var result PaymentResult

	query := `
	SELECT id, card_number, amount, currency, status, processed_at, expiry_month, expiry_year
	FROM payments
	WHERE id = ?`
	row := db.QueryRow(query, paymentID)

	err := row.Scan(&result.ID, &result.MaskedCard, &result.Amount, &result.Currency, &result.Status, &result.ProcessedAt, &result.ExpiryMonth, &result.ExpiryYear)
	if err == sql.ErrNoRows {
		return PaymentResult{}, fmt.Errorf("payment not found")
	} else if err != nil {
		return PaymentResult{}, fmt.Errorf("failed to retrieve payment: %w", err)
	}

	return result, nil
}