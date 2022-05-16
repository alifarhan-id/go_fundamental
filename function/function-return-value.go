package function

import "fmt"

func ChangeNumber(number int) (string, int) {

	number = number + 10

	return "number = ", number
}

func PrintNumber() {
	fmt.Println(ChangeNumber(60))
}
