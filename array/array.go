package array

import "fmt"

func DoArray() {
	fmt.Println("==== array =====")
	var month = [...]string{
		"january",
		"february",
		"march",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"october",
		"november",
		"desember",
	}

	for i := 0; i < len(month); i++ {

		fmt.Println("bulan ke ", i+1, "=", month[i])
	}

}
