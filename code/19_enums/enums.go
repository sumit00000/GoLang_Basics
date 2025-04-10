package main

import "fmt"

// enumerated types

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main()  {
	changeOrderStatus(Confirmed)
}