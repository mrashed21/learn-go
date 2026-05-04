// variable.go
package main

import "fmt"

func variable() {
	age := 30
	name := func() {
		user_name := "Rashed"
		fmt.Printf("User Name: %s\n Age: %d", user_name, age)
	}
	name()

}
