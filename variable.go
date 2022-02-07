package main

import "fmt"

func main() {
	var name string
	var age = 27
	var ageJhon = 35
	country := "indonesia" //declaration variable without var keyword
	var countryJhon = "america"
	name = "Jhon Doe"

	fmt.Println("==== Variable ====")

	fmt.Println("he's name is", name, ", he is", ageJhon, "years old", ", he is come from", countryJhon)

	name = "Ali Farhan"
	fmt.Println("he's name is", name, ", he is", age, "years old, ", "he is come from", country)

	//declaration multiple variable
	var (
		firstName = "jhon"
		lastName  = "doe"
	)
	fmt.Println(firstName, lastName)

}
