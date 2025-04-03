package main

import (
	"fmt"
	"time"
)

// struct -> structs are basically custon datastructures as like in other programming languages
// we craeted classes and create object and use it as like this. as in go we do not have classes
// instead of classes we use struct

type order struct {
	id string
	amount float32
	status string
	craetedAt time.Time // nanosecond precision
}

func newOrder(id string, amount float32, status string) *order  {
	// initial setup goes here...
	myOrder := order {
		id: id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

func (o *order) chnagestatus(status string)  {
	o.status = status
}

func (o order) getAmmount() float32  {
	return o.amount
}

func main()  {
	myOrder := order {
		id: "1",
		amount: 50.00,
		status: "received",
	}

	myOrder.craetedAt = time.Now()
	fmt.Println(myOrder.status)


	myOrder2 := order {
		id: "2",
		amount: 100,
		status: "delevired",
		craetedAt: time.Now(),
	}

	myOrder.chnagestatus("confirmed")
	fmt.Println(myOrder.getAmmount())


	fmt.Println("Order struct", myOrder)
	fmt.Println("Order2 struct", myOrder2)


	myOrder3 := newOrder("1", 30.50, "receiveddd")
	fmt.Println(myOrder3)

	language := struct {
		name string
		isGood bool
	} {"golang", true}

	fmt.Println("language struct", language)

}