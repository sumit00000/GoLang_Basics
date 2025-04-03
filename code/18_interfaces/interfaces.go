package main

// interfaces are basicly contract

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct{
	gateway paymenter
}
// type payment struct{
// 	gateway stripe
// }

// open close principle
// func (p payment) makepayment(amount float32) {
// 	// razerpayPaymentGW := razorpay{}
// 	stripePaymentGW := stripe{}
// 	// razerpayPaymentGW.pay(amount)
// 	stripePaymentGW.pay(amount)
// }

func (p payment) makepayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct {}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct{}

func(s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}


type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testing purpose",amount)
}


type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string ) {
	
}

func main()  {
	// stripePaymentGW := stripe{}
	// razerpayPaymentGW := razorpay{}
	paypalGW := paypal{}
	// fakeGW := fakePayment{}
	newPayment := payment{
		// gateway: stripePaymentGW,
		gateway: paypalGW,

	}
	newPayment.makepayment(100)
}