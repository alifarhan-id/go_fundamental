package function

import "fmt"

type Student struct {
	name, major, address string
}

func Mahasiswa() Student {
	var student Student
	student.name = "ali farhan"
	student.major = "electrical engineering"
	student.address = "mataram"

	return student
}

func GetMahasiswa() {
	getStudent := Mahasiswa()

	fmt.Println(getStudent.name)
}
