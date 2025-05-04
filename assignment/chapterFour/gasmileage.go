package main

import "fmt"

func main() {
	fmt.Println("Enter the miles driven for this trip below.")
	var milesDriven int
	var gallonsUsed int

	totalMiles := 0
	totalGallons := 0

	sentinelLoop := -1
	//var stopLoop int
	fmt.Println("Enter '-1' for miles driven to stop.")
	//fmt.Scanf("%d", &stopLoop)

	for {
		fmt.Println("Enter the miles driven: ")
		fmt.Scanln(&milesDriven)
		if milesDriven == sentinelLoop {
			break
		}
		fmt.Println("Enter the gallons used: ")
		fmt.Scanln(&gallonsUsed)

		if gallonsUsed == 0 {
			fmt.Println("Gallons used cannot be zero. Please enter a valid number.")
			continue
		}

		milesPerGallon := calculateMileage(milesDriven, gallonsUsed)
		fmt.Printf("Miles per gallon for this trip is: %.2f\n", milesPerGallon)

		totalMiles += milesDriven
		totalGallons += gallonsUsed

		combinedMilesPerGallonUsed := combinedMilesPerGallon(totalMiles, totalGallons)
		fmt.Printf("Combined miles per gallon is: %.2f\n", combinedMilesPerGallonUsed)
	}
	fmt.Println("Bye bye o")
}

func calculateMileage(milesDriven int, gallonUsed int) float64 {
	milePerGallon := milesDriven / gallonUsed
	return float64(milePerGallon)
}

func combinedMilesPerGallon(totalMiles int, totalGallon int) float64 {
	combinedMiles := totalMiles / totalGallon
	return float64(combinedMiles)
}
