// Write a function MySqrt which calculates the square root of a float64, but
// returns an error if this number is negative; named and unnamed return values
package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println(MySqrt(5432.0))
	fmt.Println(MySqrt2(5432.0))
}

func MySqrt(num float64) (sqrt float64, err error) {
	if num < 0 {
		return 0, errors.New("Error: input cannot be negative")
	}
	err = nil
	sqrt = num / 2.0 // first approx.
	for i := 0; i < 20; i++ {
		// refine approximation
		sqrt = (sqrt + (num / sqrt)) / 2.0
	}
	return sqrt, err

}

func MySqrt2(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Error: input cannot be negative")
	}
	err := error(nil)
	sqrt := num / 2.0
	for i := 0; i < 20; i++ {
		sqrt = (sqrt + (num / sqrt)) / 2.0
	}
	return sqrt, err

}
