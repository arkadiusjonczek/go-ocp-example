package main

type PaymentProvider interface {
	pay(amount float64) error
}

type PaymentServiceInterface interface {
	pay(amount float64) error
}

var _ PaymentServiceInterface = (*PaymentService)(nil)

type PaymentService struct {
	provider PaymentProvider
}

func NewPaymentService(provider PaymentProvider) PaymentServiceInterface {
	return &PaymentService{
		provider: provider,
	}
}

func (service *PaymentService) pay(amount float64) error {
	// do other relevant stuff
	err := service.provider.pay(amount)
	// do other relevant stuff

	return err
}
