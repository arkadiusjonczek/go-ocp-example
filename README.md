# go-ocp-example

**Example of the Open Closed Principle (OCP) from SOLID.**

The `PaymentService` implements the high level abstraction of the payment for example in a online shop.

The `PaymentProvider` interface can be implemented by concrete payments providers.

Adding a new payment provider requires to create a new class implementing the `PaymentProvider` interface. But it's not required to change the `PaymentService` for that.