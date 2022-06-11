package main

import "fmt"

func main() {
	// declaring variables
	var str string
	var a, b int
	var ab float32
	// assigning variables
	str = "Hello World"
	a = 50
	b = 25
	ab = 36.88

	fmt.Println("The value str=", str)
	fmt.Println("The value a=",a)
	fmt.Println("The value b=",b)
	fmt.Println("The value ab=",ab)

	// declare and assign values to variables
	var city string = "India"
	var x int = 20
	fmt.Println("the name of city=", city)
	fmt.Println("the number is =", x)

	// declare variable with  defining its types
	country:= "India"
	val:= 25
	fmt.Println("the country name =", country)
	fmt.Println("the country val =", val)


	// defining multiple variables
	var  (
		name string
		email string
		age int
	)
	name = "Asd"
	email ="asd@gmail.com"
	age = 50
	
	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(age)

}