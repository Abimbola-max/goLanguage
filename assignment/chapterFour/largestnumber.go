package main

import "fmt"

//func largestNumber() {
//	var inputs int
//	fmt.Println("Enter 10 different numbers")
//	fmt.Scanf("%d", &inputs)
//
//}

func main() {
	largest := 0
	var num [10]int
	for i := 0; i < len(num); i++ {
		fmt.Println("Enter each number: ")
		fmt.Scanf("%d", &num[i])
		if num[i] > largest {
			largest = num[i]
		}
	}
	fmt.Printf("The largest number is out of the 10 inputs is: %d", largest)
}
