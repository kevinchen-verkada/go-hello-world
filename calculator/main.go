package main

import (
	"fmt"
	"log"

	"golang.kevinchen-verkada/calculator/mathutils"
)

func main() {
	// Create a new calculator
	calc := mathutils.NewCalculator()
	fmt.Printf("Created calculator: %s\n\n", calc)

	// Perform some calculations
	fmt.Println("Performing calculations...")

	result := calc.Add(10, 5)
	fmt.Printf("10 + 5 = %.2f\n", result)

	result = calc.Multiply(result, 2)
	fmt.Printf("%.2f * 2 = %.2f\n", 15.0, result)

	result, err := calc.Divide(result, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f / 3 = %.2f\n", 30.0, result)

	result = calc.Power(2, 8)
	fmt.Printf("2^8 = %.2f\n", result)

	// Show calculation history
	fmt.Println("\nCalculation History:")
	history := calc.GetHistory()
	for i, op := range history {
		fmt.Printf("%d. %s(%.2f, %.2f) = %.2f at %s\n",
			i+1, op.Type, op.Arguments[0], op.Arguments[1],
			op.Result, op.Timestamp.Format("15:04:05"))
	}

	// Get last result
	lastResult, err := calc.GetLastResult()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nLast result: %.2f\n", lastResult)

	// Demonstrate error handling
	fmt.Println("\nTesting division by zero...")
	_, err = calc.Divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
