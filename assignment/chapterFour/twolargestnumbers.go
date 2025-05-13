//package main
//
//func main() {
//	numbers := [10]int{12, 34, 45, 50, 30, 5, 45, 4, 1, 80}
//	firstLargest := findFirstLargestNumber(numbers)
//}
//
//func findFirstLargestNumber(numbers []int) int {
//	largest := 0
//	//var numbers [10]int
//	for i := 0; i < len(numbers); i++ {
//		//fmt.Println("Enter each number: ")
//		//fmt.Scanf("%d", &num[i])
//		if numbers[i] > largest {
//			largest = numbers[i]
//		}
//	}
//	return largest
//}
//
//func findSecondLargestNumber(numbers []int) int {
//	largest := findFirstLargestNumber(numbers)
//	//var num [10]int
//	for i := 0; i < len(numbers); i++ {
//		if numbers[i] > largest {
//			largest = numbers[i]
//		}
//	}
//	return largest
//}

package main // Every Go program must start with a package declaration [1]

import "fmt" // Import the format package for input/output operations [2]

func findFirstLargestNumber(numbers []int) int {
	// Variables are storage locations with a name and type [5].
	// We initialize 'largest' with the first element, assuming the slice is not empty.
	// For robustness in a real-world scenario, you might handle an empty slice [6].
	// Given the example array has a fixed size of 10, the slice will not be empty.
	if len(numbers) == 0 { // len() returns the length of a slice [7]
		// Returning 0 might be misleading if 0 is a valid input number.
		// For this example based on non-negative numbers, it might suffice,
		// but returning an error or a sentinel value (like math.MinInt from an external package)
		// or checking the return value in main is better practice [8].
		return 0 // Simplified return for this example
	}

	largest := numbers // Use short variable declaration [9], initialize with the first element [10]

	// Use a for loop with the range keyword to iterate over the elements of the slice [11, 12].
	// The underscore _ is used to ignore the index variable returned by range, as we only need the value [13].
	for _, number := range numbers {
		// Use an if statement to compare the current number with the largest found so far [14, 15].
		if number > largest { // The > operator is used for comparison [16]
			largest = number // Assign the current number as the new largest if it's greater [17]
		}
	}
	return largest // Return the final largest number [18]
}

// findSecondLargestNumber finds the second largest distinct number in a slice.
// It takes the slice ([]int) and the first largest number (int) as input parameters [3, 4].
// It returns the second largest distinct integer found (int).
// It handles cases where there are fewer than two distinct numbers by returning an initial value (-1).
func findSecondLargestNumber(numbers []int, firstLargest int) int {
	// Initialize 'secondLargest'. We need a value that is smaller than any possible
	// valid second largest number. Assuming the input numbers are likely non-negative
	// based on your example, initializing to -1 is a common way to indicate "not found yet"
	// and will be smaller than any non-negative number.
	secondLargest := -1 // Initialize secondLargest to a value indicating it hasn't been found

	// Iterate through the numbers in the slice using a for loop with range [11, 12].
	for _, number := range numbers { // Iterate through the values [12, 13]
		// We need to check if the current number is a candidate for the second largest.
		// It must satisfy two conditions:
		// 1. It must be greater than the current 'secondLargest' candidate.
		// 2. It must *not* be equal to the 'firstLargest' number already found.
		if number > secondLargest && number != firstLargest { // Use comparison operators [16] and logical AND [19, 20]
			secondLargest = number // If both conditions are true, this number is the new second largest candidate [17]
		}
		// This logic correctly finds the largest number among those not equal to 'firstLargest'.
	}

	return secondLargest // Return the found second largest number [18] (or -1 if no distinct second largest was found > -1)
}

// The main function is the entry point of the program [21].
func main() {
	// Initialize an array with the given numbers [22].
	// Arrays have a fixed length, specified inside the brackets [22].
	numbersArray := [23]int{12, 34, 45, 50, 30, 5, 45, 4, 1, 80} // Declaring and initializing an array

	// Functions can operate on slices. We can create a slice that refers to the entire array.
	// Using array[:] creates a slice covering the whole array [24].
	numbersSlice := numbersArray[:] // Create a slice from the array

	// Print the list of numbers we are processing for clarity [25, 26].
	fmt.Println("Processing the list of numbers:", numbersSlice) // Print using fmt.Println

	// Call the findFirstLargestNumber function, passing the slice as an argument [3].
	// Assign the returned value to the 'firstLargest' variable using short declaration [9, 18].
	firstLargest := findFirstLargestNumber(numbersSlice)

	// Call the findSecondLargestNumber function, passing the slice and the first largest number as arguments [3].
	// Assign the returned value to the 'secondLargest' variable [9, 18].
	secondLargest := findSecondLargestNumber(numbersSlice, firstLargest)

	// Display the results using fmt.Println [25, 26].
	fmt.Println("The first largest number is:", firstLargest)

	// Check if a valid distinct second largest number was found (i.e., its value is not the initial -1
	// and it's not equal to the first largest, which could happen if all numbers were the same and positive).
	if secondLargest != -1 && secondLargest != firstLargest { // Use an if statement for conditional output [14, 15]
		fmt.Println("The second largest number is:", secondLargest) // Print the second largest number
	} else {
		// Handle cases where there isn't a distinct second largest number
		// (e.g., array has only one unique value, or is empty/too short, though we assumed non-empty).
		fmt.Println("There is no distinct second largest number in the list.")
	}

	fmt.Println("Finished finding the largest numbers.") // End message
}
