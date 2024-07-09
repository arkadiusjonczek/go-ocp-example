package main

import "log"

var _ PaymentProvider = (*PayPal)(nil)

type PayPal struct {
}

func NewPayPal() *PayPal {
	return &PayPal{}
}

func (provider *PayPal) pay(amount float64) error {
	// paypal payment logic
	log.Println("PayPal Payment")
	return nil
}

var _ PaymentProvider = (*CreditCard)(nil)

type CreditCard struct {
}

func NewCreditCard() *CreditCard {
	return &CreditCard{}
}

func (provider *CreditCard) pay(amount float64) error {
	// credit card payment logic
	log.Println("CreditCard Payment")
	return nil
}
