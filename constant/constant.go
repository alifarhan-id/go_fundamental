package constant

import "fmt"

func DoConstant() {
	const firstName = "ali" //declaration constant
	const lastName = "farhan"

	fmt.Println(firstName)

	const (
		word   = "eat"
		number = 30
	)

	fmt.Println(word, number)
}
