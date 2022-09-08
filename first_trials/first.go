package first_trials

import (
	"fmt"
	"sync"
	"time"
)

var (
	a     int
	insan string
	b     float32
)

type (
	mystr struct {
		name    string
		surname string
		age     int
	}

	mytype int
)

var concurr = sync.WaitGroup{}

func main() {

	// fmt.Println(newj["a"])
	// for key, val := range newj {
	// 	fmt.Printf("key = %v, value = %v\n", key, val)
	// }
	// myf(newj["a"], newj["b"], newj["c"], 15, 22)

	//
	var z mystr
	z.name = "hamza"
	z.surname = "kaya"
	z.age = 20

	concurr.Add(2)
	go print(z.age)
	go z.prt()
	concurr.Wait()

}
func print(age int) int {
	for i := 0; i < 10; i++ {
		fmt.Println(age)
		age++
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	concurr.Done()
	return 0
}

func (u *mytype) point() {
	fmt.Println("thats my pointer method bro", (*u))
}

func (z mytype) ascii() {
	fmt.Printf("ascii value of the number %d is %c", z, z)
}

func (z mytype) hello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello", z)
		z++
		time.Sleep(time.Duration(5 * time.Millisecond))

	}
	concurr.Done()

}

func (x mystr) prt() {
	fmt.Println(x.name)
	time.Sleep(time.Duration(10 * time.Millisecond))
	fmt.Println(x.surname)
	time.Sleep(time.Duration(10 * time.Millisecond))
	fmt.Println(x.age)
	time.Sleep(time.Duration(10 * time.Millisecond))
	concurr.Done()
}

// func myf(numbers ...int) (a int) {
// 	fmt.Println(len(numbers))
// 	for index, value := range numbers {
// 		fmt.Println("index = ", index, "value = ", value)
// 	}
// 	// a = numbers[0]
// 	// fmt.Println(a)
// 	return

// }
