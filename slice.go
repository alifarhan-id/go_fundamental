package main

import "fmt"

func main() {
	fmt.Println("====== slice ===")

	month := [...]string{
		"january",
		"february",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	slice1 := month[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	month[5] = "change"
	fmt.Println(slice1)

}
