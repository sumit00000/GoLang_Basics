package main

import (
	"fmt"
	"time"
)


type customer struct {
	name string
	phone string
}

type order struct {
	id string
	amount float32
	status string
	craetedAt time.Time  // nanosecond precision
	customer                                 
}

func main()  {
	newcustomer := customer {
		name: "sumit",
		phone: "98766568",
	}
	newOrder := order {
		id: "1",
		amount: 30,
		status: "received",
		customer: newcustomer,
	}

	// newOrder := order {
	// 	id: "1",
	// 	amount: 30,
	// 	status: "received",
	// 	customer: customer{
	// 		name: "sumit",
	// 		phone: "98766568",
	// 	},
	// }

	// change name
	newOrder.customer.name = "singh"

	fmt.Println(newOrder)
}