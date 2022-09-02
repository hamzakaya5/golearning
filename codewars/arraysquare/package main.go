// package main

// import (
// 	"fmt"
// 	"math"
// )

// func squareOrSquareRoot(arr [5]int) []int {
// 	z := make([]int, 5)
// 	for i := 0; i < len(arr); i++ {

// 		x := math.Sqrt(float64(arr[i]))
// 		if math.Pow(x, 2) == float64(arr[i]) {
// 			z[i] = int(math.Sqrt(float64(arr[i])))
// 		} else {
// 			z[i] = int(math.Pow(float64(arr[i]), float64(2)))
// 		}
// 	}
// 	fmt.Println(z)
// 	return z
// }

// func main() {
// 	var myarr [5]int = [5]int{4, 5, 8, 9, 49}
// 	squareOrSquareRoot(myarr)
// }

package main

var GlobalFlag string

func parantheses(x string) (y bool) {
	y = true
	idx := 0
	for idx = 0; idx != len(x)/2; idx++ {
		if x[idx] == x[len(x)-idx-1] {
			y = false
			break
		}
	}
	return
}

func main() {
	GlobalFlag = "(()(()()()))"
	z := parantheses(GlobalFlag)
	print(z)
}
