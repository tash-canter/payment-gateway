package services

func mockBankSimulation() (string, string){
    paymentID := generatePaymentID()
    return paymentID, "Success"
}