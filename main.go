package main

func main() {
	amount := float64(1337)

	payPalPaymentService := NewPaymentService(NewPayPal())
	payPalPaymentService.pay(amount)

	creditCardPaymentService := NewPaymentService(NewCreditCard())
	creditCardPaymentService.pay(amount)
}
