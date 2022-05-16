package function

import "fmt"

func Bio(name string, age int, address string) (string, int, string) {
	return name, age, address
}

func CallBio() {
	out1, _, _ := Bio("ali farhan", 27, "west nusa tenggara")
	fmt.Println(out1)
}
