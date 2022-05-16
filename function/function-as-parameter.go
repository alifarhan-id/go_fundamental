package function

import "fmt"

type Filter func(string) string

func StringWithFilter(words string, filter Filter) {
	wordsWithFilter := filter(words)

	fmt.Println(wordsWithFilter)

}

func Check(words string) string {
	if words == "Anjing" || words == "anjing" {
		return "*****"
	} else {
		return words
	}
}

func CallStringWithFilter() {
	filter := Check

	StringWithFilter("ali farhan", filter)
	StringWithFilter("anjing", filter)
	StringWithFilter("Anjing", filter)
}
