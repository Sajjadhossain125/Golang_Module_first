package main

import "fmt"

/*
The Printf() Function
The Printf() function first formats its argument based on the given formatting verb and then prints them.

Here we will use two formatting verbs:

%v is used to print the value of the arguments
%T is used to print the type of the arguments*/
const address string = "house-6,StarkulRoad,Dhaka"
const age int = 25
const phoneNumb string = "01908064229"
const skillmark float32 = 32.5

func printfOutput() {

	fmt.Printf("My home address: %v type %T\n", address, address)
	fmt.Printf("I'm %v type %T\n", age, age)
	fmt.Printf("my mobile number %v type %T\n", phoneNumb, phoneNumb)
	fmt.Printf("my skill level: %v type %T\n", skillmark, skillmark)

}
