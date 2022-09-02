package error

import (
	"fmt"
)

func OsErr() {
	// 	f, err := os.Open("myerror.txt")
	// 	if err != nil {
	// 		if pErr, ok := err.(*os.PathError); ok {
	// 			fmt.Println("file could not found my bro", pErr.Path)
	// 			return
	// 		}
	// 	}

}

type element struct {
	name    string
	surname string
	age     int
}

func (kisi element) ageCheck() {
	if kisi.age < 18 {
		fmt.Println(kisi.name, kisi.surname, "Girisiniz yasaklanmistir")
	} else {
		fmt.Println(kisi.name, kisi.surname, "Geçebilirsiniz")
	}
}

func Myinter(a interface{}) {
	x, ok := a.(int)
	if ok {
		fmt.Println(x)
	} else {
		fmt.Println("tyep is not integer")
	}
}

func AgeCheck() {
	people := element{name: "hamza", surname: "kaya", age: 18}
	people.ageCheck()
}
