package main

import "fmt"

const baseWeeklySalary float64 = 200.00
const commissionRate float64 = 0.09

func getInputSales() float64 {
	var totalGrossSales float64 = 0.0

	fmt.Println("Sales Commission Calculator")
	fmt.Println("Enter item values sold one by one.")
	fmt.Println("Enter 0 or a negative number to finish input.")

	for {
		var itemValue float64
		fmt.Print("Enter item value: $")
		fmt.Scanln(&itemValue)

		if itemValue <= 0 {
			break
		}
		totalGrossSales += itemValue
	}

	return totalGrossSales
}

func calculateEarnings(totalSales float64) float64 {

	commission := totalSales * commissionRate

	totalWeeklyEarnings := baseWeeklySalary + commission

	return totalWeeklyEarnings
}

func main() {
	salesMade := getInputSales()

	weeklyEarnings := calculateEarnings(salesMade)

	fmt.Println("\n--- Calculation Summary ---")
	fmt.Printf("Base Salary: $%.2f\n", baseWeeklySalary)
	fmt.Printf("Commission Rate: %.2f%%\n", commissionRate*100)
	fmt.Printf("Total Gross Sales: $%.2f\n", salesMade)
	fmt.Printf("Commission Earned: $%.2f\n", salesMade*commissionRate)
	fmt.Printf("Total Weekly Earnings: $%.2f\n", weeklyEarnings)
}
