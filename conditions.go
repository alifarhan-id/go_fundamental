package main

import "fmt"

func main() {
	point := 10

	if point == 10 {
		fmt.Println("pass with perfect score")
	} else if point > 5 {
		fmt.Println("passed")
	} else if point == 4 {
		fmt.Println("almost passed")
	} else {
		fmt.Println("not passed, your score is", point)
	}

	//variable temporary
	value := 8840.0

	if percent := value / 100; percent >= 100 { //percent is variable temporary
		fmt.Printf("%.1f%s perfect!\n", percent, " %")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good!\n", percent, " %")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, " %")
	}

	var number = 6

	switch number {
	case 8:
		fmt.Println("perfect")
	case 6:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}
