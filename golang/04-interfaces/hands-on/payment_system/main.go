package main

import "fmt"

// --- 1. The Strategy Interface ---
type PaymentProcessor interface {
	Pay(amount float64) string
}

// --- 2. Concrete Strategies ---

type CreditCard struct {
	CardNumber string
}

func (c CreditCard) Pay(amount float64) string {
	// Logic to charge card...
	return fmt.Sprintf("Paid $%.2f using Card ending in %s", amount, c.CardNumber[12:])
}

type PayPal struct {
	Email string
}

func (p PayPal) Pay(amount float64) string {
	// Logic to call PayPal API...
	return fmt.Sprintf("Paid $%.2f from PayPal account %s", amount, p.Email)
}

type Bitcoin struct {
	WalletAddress string
}

func (b Bitcoin) Pay(amount float64) string {
	return fmt.Sprintf("Sent $%.2f equivalent to wallet %s", amount, b.WalletAddress)
}

// --- 3. Context (The Store) ---
type checkoutService struct {
	processor PaymentProcessor // Dependency Injection
}

func (s *checkoutService) Checkout(amount float64) {
	// Delegate work to the interface
	receipt := s.processor.Pay(amount)
	fmt.Println("Receipt:", receipt)
}

func main() {
	// Setup strategies
	card := CreditCard{CardNumber: "1234567812349999"}
	paypal := PayPal{Email: "user@example.com"}

	// Dependency Injection in action

	fmt.Println("--- Customer 1 (Credit Card) ---")
	service := checkoutService{processor: card}
	service.Checkout(100.50)

	fmt.Println("--- Customer 2 (PayPal) ---")
	service.processor = paypal // Hot-swap strategy
	service.Checkout(25.00)
}
