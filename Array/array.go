package main

import "fmt"

func main() {
	//functions call
	basicArray()
	dynamicArray()
	lenghtofArray()
	stringofArray()
	modificationofArray()

	//now show the length of array

}
func basicArray() {

	//now i will show the array example of go lang

	var numbers = [3]int{1, 2, 3}
	secondlist := [5]int{4, 5, 6, 7, 8}
	fmt.Println(numbers)
	fmt.Println(secondlist)
}

//dynamic array declare
func dynamicArray() {
	var num = [...]int{11, 20, 40}
	num2 := [...]int{1: 43, 4: 50, 6: 60}
	fmt.Println(num)
	fmt.Println(num2)
}

func lenghtofArray() {
	num2 := [...]int{1: 43, 4: 50, 6: 60}
	fmt.Println(len(num2))
}

func stringofArray() {

	var foods = [5]string{"Apple", "Lemon", "Mango", "Banana", "jackfruits"}
	fmt.Printf("This is list of food, The items of foods are: %v\n", foods)
}

func modificationofArray() {

	cars := [...]string{"BMW", "Toyota", "Lamborgrini", "Porsi", "Mercities", "Bugerties", "Chaborlent"}

	fmt.Printf("The most expencibe cares are:\n %v\n,%v\n,%v\n,%v\n", cars[0], cars[2], cars[3], cars[5])
	//change the value of array "Porsi" and add "Hundai"

	cars[3] = "Hundai"

	fmt.Println(cars)

}
