package romanconvert

import (
	"fmt"
)

var (
	s       [10]string
	counter int32 = 0
)

func ConvertToRomanNumber(num int32) []string {

	// Write your code here

	if num%1000 != 0 {
		a := num / 1000
		for x := int32(0); x < a; x++ {
			fmt.Printf("M")
			s[counter] = "M"
			counter++
		}
		num = num % 1000
	}
	if num%500 != 0 {
		a := num / 500
		for x := int32(0); x < a; x++ {
			fmt.Printf("D")
			s[counter] = "D"
			counter++
		}
		num = num % 500
	}
	if num%100 != 0 {
		a := num / 100
		for x := int32(0); x < a; x++ {
			fmt.Printf("C")
			s[counter] = "C"
			counter++
		}
		num = num % 100
	}
	if num%50 != 0 {
		a := num / 50
		for x := int32(0); x < a; x++ {
			fmt.Printf("L")
			s[counter] = "L"
			counter++
		}
		num = num % 50
	}
	if num%10 != 0 {
		a := num / 10
		for x := int32(0); x < a; x++ {
			fmt.Printf("X")
			s[counter] = "X"
			counter++
		}
		num = num % 10
	}
	if num%5 != 0 {
		a := num / 5
		for x := int32(0); x < a; x++ {
			fmt.Printf("v")
			s[counter] = "V"
			counter++
		}
		num = num % 5
	}
	if num%4 != 0 {
		a := num / 4
		for x := int32(0); x < a; x++ {
			fmt.Printf("Iv")
			s[counter] = "IV"
			counter++
		}
		a = num % 4
		for x := int32(0); x < a; x++ {
			fmt.Printf("I")
			s[counter] = "I"
			counter++
		}

	}
	retVal := s[:]
	return retVal
}
