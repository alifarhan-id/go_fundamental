package main

import (
	"alifarhan1230/go_fundamental/Map"
	"alifarhan1230/go_fundamental/array"
	"alifarhan1230/go_fundamental/boolean"
	"alifarhan1230/go_fundamental/conditions"
	"alifarhan1230/go_fundamental/constant"
	"alifarhan1230/go_fundamental/conversion"
	"alifarhan1230/go_fundamental/helloworld"
	"alifarhan1230/go_fundamental/number"
	"alifarhan1230/go_fundamental/operator"
	"alifarhan1230/go_fundamental/slice"
	"alifarhan1230/go_fundamental/string"
	"alifarhan1230/go_fundamental/variable"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("============== Welcome to Basic learning GO ====================")
	fmt.Println("================= Choose menu what you want... =======================")

	fmt.Println("1. Hello World")
	fmt.Println("2. Number")
	fmt.Println("3. Boolean")
	fmt.Println("4. String")
	fmt.Println("5. Variable")
	fmt.Println("6. Constant")
	fmt.Println("7. Conversion")
	fmt.Println("8. Type Declaration")
	fmt.Println("9. operation")
	fmt.Println("10. conditions")
	fmt.Println("11. Array")
	fmt.Println("12. Slice")
	fmt.Println("13. Map")
	fmt.Println("99. Exit")

	for {
		doAction()
	}

}

func doAction() {
	fmt.Print("MENU : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	menu, err := strconv.ParseInt(scanner.Text(), 10, 64)

	if menu < 1 || err != nil {
		fmt.Println("invalid input")
		return
	}

	switch menu {
	case 1:
		helloworld.DoHelloWorld()
	case 2:
		number.DoNumber()
	case 3:
		boolean.DoBoolean()
	case 4:
		string.DoString()
	case 5:
		variable.DoVariable()
	case 6:
		constant.DoConstant()
	case 7:
		conversion.DoConversion()
	case 8:
		conversion.DoConversion()
	case 9:
		operator.DoOperator()
	case 10:
		conditions.DoConditions()
	case 11:
		array.DoArray()
	case 12:
		slice.DoSlice()
	case 13:
		Map.DoMap()

	case 99:
		os.Exit(2)

	default:

	}
}
