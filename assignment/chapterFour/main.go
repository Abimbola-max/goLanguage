package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	x := 5
	fmt.Println(x)
	var name string = "Abimbola"
	fmt.Printf("my name is %s\n", name)
	//for i := 1; i <= 10; i++ {
	//	fmt.Println(i)
	//}
	var numbers [4]float64
	numbers[0] = 20
	numbers[1] = 10
	numbers[2] = 15
	numbers[3] = 7
	var total float64 = 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	fmt.Println(total)
	fmt.Println(total / 5)
}
