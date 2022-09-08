package main

// func printArr(x []int) {
// 	for idx, elemnt := range x {
// 		print(idx, " ", elemnt, "\n")
// 	}

// }

// // addOdd takes the array of sorted odd numbers and changes the odd element of the array 'x'
// func addOdd(x []int, y []int) {
// 	counter := 0
// 	for i := 0; i < len(x); i++ {
// 		if x[i]%2 != 0 {
// 			x[i] = y[counter]
// 			counter++
// 		}
// 	}
// }

// //sortOdd takes the array and find its odd elements and sort those elements
// func sortOdd(x []int) {
// 	var index []int
// 	for _, elmnt := range x {
// 		if elmnt%2 != 0 {
// 			index = append(index, elmnt)
// 		}
// 	}

// 	for i := 0; i < len(index); i++ {
// 		for j := 0; j < i; j++ {
// 			if index[i] < index[j] {
// 				temp := index[i]
// 				index[i] = index[j]
// 				index[j] = temp
// 			}
// 		}
// 	}
// 	addOdd(x, index)
// }

// func main() {
// 	var arr = [...]int{5, 8, 2, 4, 1, 53, 54, 90, 6, 9, 7, 43, 12, 87, 91, 3, 10}
// 	sortOdd(arr[:])
// 	printArr(arr[:])
// }

//////////////////////////////////////////////////////
// func RemainderSorting(strArr []string) []string {
//     for j:= 0; j < len(strArr); j++ {

//         for i:= 0; strArr[i+1] != nil; i++{
//             if len(strArr[i]%3 > len(strArr[i+1]%3) {
//                 temp:= strArr[i]
//                 strArr[i] = strArr[i+1]
//                 strArr[i+1] = temp
//             }else if len(strArr[i]%3 == len(strArr[i+1]%3){
//                 slice[i:i+2]
//             }
//         }
//     }
// }

/////////////////////////////////////
// func TwoSum(numbers []int, target int) [2]int {
// 	i := 0
// 	j := 0
// 	breakPoint := 0
// 	var arr [2]int
// 	for i = 0; i < len(numbers); i++ {
// 		fmt.Println("inside for loop")
// 		for j = 0; j < i; j++ {
// 			if numbers[i]+numbers[j] == target {
// 				arr[0] = j
// 				arr[1] = i
// 				breakPoint = 1
// 				break
// 			}
// 		}
// 		if breakPoint == 1 {
// 			break
// 		}
// 	}
// 	return arr
// }

// func main() {
// 	x := [5]int{10, 20, 30, 40, 50}
// 	target := 90
// 	a := TwoSum(x[:], target)
// 	fmt.Println(a)
// }
