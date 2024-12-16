package services

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func maskCardNumber(cardNumber string) string {
	if len(cardNumber) < 4 {
		return "****"
	}
	return "**** **** **** " + cardNumber[len(cardNumber)-4:]
}

func validatePaymentRequest(req PaymentRequest) error {
	if req.Amount <= 0 {
        return fmt.Errorf("invalid payment amount: must be greater than 0")
    }
    if req.Currency == "" {
        return fmt.Errorf("currency must be specified")
    }
	if !isValidCardNumber(req.CardNumber) {
        return fmt.Errorf("invalid card number")
    }
    if !isValidCVV(req.CVV) {
        return fmt.Errorf("invalid CVV")
    }
    if !isValidExpiryDate(req.ExpiryMonth, req.ExpiryYear) {
        return fmt.Errorf("card has expired")
    }
	return nil
}

func isValidCardNumber(cardNumber string) bool {
	// Luhn's algorithm
    var sum int
	shouldDouble := false
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')
		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		shouldDouble = !shouldDouble
	}
	return sum%10 == 0
}

func isValidCVV(cvv string) bool {
    if len(cvv) < 3 || len(cvv) > 4 {
        return false
    }
    for _, digit := range cvv {
        if digit < '0' || digit > '9' {
            return false
        }
    }
    return true
}

func isValidExpiryDate(month int, year int) bool {
    if month < 1 || month > 12 {
        return false
    }

    currentYear, currentMonth := time.Now().Year(), int(time.Now().Month())
    if year > currentYear {
        return true
    }

	if year == currentYear && month >= currentMonth {
        return true
    }

    return false
}

func generatePaymentID() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	return fmt.Sprintf("%x", rand.Int63())
}