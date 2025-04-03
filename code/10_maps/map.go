package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main()  {
	// creating map
	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// get an element
	fmt.Println(m["name"])
	fmt.Println(m["area"])

	// IMP: if key value does not exist in the map then it returns zero
	fmt.Println(m["line"])


	n := make(map[string]int)
	n["age"] = 30
	n["price"] = 50
	fmt.Println(n["age"])
	fmt.Println(n["phone"])
	fmt.Println(len(n))

    delete(n,"price")
	fmt.Println(n)



	l := map[string]int{"price": 40, "phone": 3}
	fmt.Println(l)
	v, ok := l["price"]     // v means it return value and ok mean true or false
	fmt.Println(v)

	if ok {
		fmt.Println("All Ok!")
	} else {
		fmt.Println("not ok!")
	}



    
    // map equality us emaps package
	m1 := map[string]int{"phone": 30, "cable": 50}
	m2 := map[string]int{"phone": 30, "cable": 50}

	fmt.Println(maps.Equal(m1, m2))


}