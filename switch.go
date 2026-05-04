package main

import "fmt"

func switchCase() {
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday.")
	case "Tuesday":
		fmt.Println("Today is Tuesday.")
	case "Wednesday":
		fmt.Println("Today is Wednesday.")
	default:
		fmt.Println("It's another day.")
	}
}
