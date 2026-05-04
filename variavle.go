// variable.go
package main

import "fmt"

func variable() {
	age := 30 // outer variable
	name := func() {
		user_name := "Rashed" // inner variable
		fmt.Printf("User Name: %s\n Age: %d", user_name, age)
	}
	name()

	/* outer inner e access korte parbe but inner theke outer access kora jabe na.*/

}
