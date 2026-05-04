// variable.go
package main

import "fmt"

func variable() {
	age := 30 // outer variable
	name := func() {
		user_name := "Rashed" // inner variable
		fmt.Printf("User Name: %s\n Age: %d\n", user_name, age)
	}
	name()

	/* outer inner e access korte parbe but inner theke outer access kora jabe na.*/

	// variable shadowing
	Age := 25
	shadowing := func() {
		Age = 20                                // this age variable shadows the outer age variable
		fmt.Printf("\nShadowed Age: %d\n", Age) // prints 20
	}
	shadowing()
	fmt.Printf("Outer Age: %d\n", Age) // prints 20
}
