package function

import "fmt"

func GetDoCount() {
	DoCount()
}

func DoCount() {
	var (
		number1 = 10
		number2 = 30
	)

	result := number1 + number2
	fmt.Println(result)

}
