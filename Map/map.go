package Map

import "fmt"

func DoMap() {
	var book map[string]string = make(map[string]string)
	book["tittle"] = "learn Go-Lang"
	book["author"] = "invoker"

	fmt.Println(book)

}
