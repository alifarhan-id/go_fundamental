package array

import "fmt"

func DoArray() {
	fmt.Println("==== array =====")
	var vocab = [6]string{
		"eat",
		"sleep",
		"admire",
		"bring",
		"dring",
		"love",
	}

	for i := 0; i < len(vocab); i++ {
		fmt.Println(vocab[i])
		if i == 6 {
			fmt.Println("array len is 6")
		}
	}
}
