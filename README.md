# Building a Payment Gateway

To run my solution:

- Install go and dependencies: `go mod tidy`
- To run the application use the following command: `go run main.go`
  This will start the application on port 8080 and automatically create the SQLite database if it doesn't already exist. The app listens for HTTP requests and simulates a payment gateway system.
- To test the API:
  - Make a POST request to /payments with the necessary card details (card number, expiry month and year, CVV, amount, currency).
    Example JSON payload:
    ```json
    {
      "card_number": "4539148803436467",
      "expiry_month": 12,
      "expiry_year": 2025,
      "cvv": "123",
      "amount": 100.5,
      "currency": "USD"
    }
    ```
  - Make a GET request to /payments/{id}, where {id} is the unique payment ID returned when the payment was processed.

Assumption:

- SQLite is used for simplicity here. However, it only supports one write operation at a time, which could limit performance in a high-concurrency environment.
- The payment gateway is mocked based on a simple check to see if the card number passes the Luhn algorithm check.

Improvements:

- In a production environment a server based database like PostgreSQL would be better to handle higher traffic and concurrency.
- Adding more complex bank simulation such as further validation to the card and other conditions so that it isn't always success. E.g. timeout responses or fail if there is insufficient funds.
- Adding more security measures to ensure the confidentiality of payment details
- Adding more tests:
  - Unit tests to ensure each function in the system works as expected
  - Edge case testing, e.g. invalid card numbers, expiry dates, invalid or missing fields

Cloud technologies:

- Docker could be used for containerising the app for easier deployment
- Database hosted on cloud provider such as AWS or Google Cloud
