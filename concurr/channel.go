package concurr

import (
	"fmt"
	"time"
)

func Channel1(a string, mychan chan int) {
	x := len(a) % 2
	if x == 0 {
		time.Sleep(1 * time.Second)
		fmt.Println("the string length is even")
	} else {
		fmt.Println("the string length is odd")
	}
	mychan <- x
}

func Channel2(a string, mychan chan int) {
	x := len(a)
	if x == 0 {
		fmt.Println("the string length is even")
	} else {
		fmt.Println("the string length is odd")
	}
	mychan <- x
}
