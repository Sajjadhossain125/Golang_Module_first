package main

import "fmt"

func multipleConst() {
	const (
		name     string = "mdsajjad"
		age      int    = 25
		marriade bool   = false
	)
	fmt.Println("this information is mine", name, age, marriade)
}
