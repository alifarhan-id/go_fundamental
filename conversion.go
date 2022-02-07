package main

import "fmt"

func main() {
	var value32 int32 = 100000
	var value64 int64 = int64(value32)
	var value8 int8 = int8(value32)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value8) //why -96 ? coz is happend value of datatype overflow is mean when value beyond the limit the value directly back to minimum value

	name := "jhon"
	e := name[1]
	eStirng := string(e)

	fmt.Println(name)
	fmt.Println(eStirng)
}
