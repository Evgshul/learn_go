package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64, initialGuess float64) float64 {

	z := initialGuess
	threshold := 1e-10
	iterations := 0

	for {
		iterations++
		preZ := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-preZ) < threshold {
			break
		}
	}
	fmt.Printf("Stable result reached in %d iterations.\n", iterations)
	return z
}

func main() {
	valuesToSquare := []float64{1, 2, 3, 4, 25, 90}
	initialGuess := []float64{1.0, 0.5, 2}
	for _, guess := range initialGuess {
		fmt.Printf("Initial guess: %v\n", guess)
		for _, x := range valuesToSquare {
			result := Sqrt(x, guess)
			realResult := math.Sqrt(x)
			difference := math.Abs(realResult - result)

			fmt.Printf("x = %v, ourResult = %v, Real Result = %v, difference = %v\n", x, result, realResult, difference)
			if difference == 0 {
				break
			}
		}
		fmt.Println()
	}
	// fmt.Println(Sqrt(2))
}
