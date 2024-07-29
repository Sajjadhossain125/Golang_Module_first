package main

import (
	"fmt"
)

/*
Verb	Description
%v	Prints the value in the default format
%#v	Prints the value in Go-syntax format
%T	Prints the type of the value
%%	Prints the % sign
*/
func main() {
	const name = "my name"
	const data = 5
	fmt.Printf("%v \n", name)
	fmt.Printf("%#v \n", name)
	fmt.Printf("%T \n", name)
	fmt.Printf("%v%% \n", name)
	fmt.Println("this is the next section")
	fmt.Printf("%v \n", data)
	fmt.Printf("%#v \n", data)
	fmt.Printf("%T \n", data)
	fmt.Printf("%v%% \n", data)

}
