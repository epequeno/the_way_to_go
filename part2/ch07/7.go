/* a) Write a function Sum which has as parameter an array arrF of 4 floating
point numbers, and which returns the sum of all the numbers in the array

How would the code have to be modified to use a slice instead of an array?

Answer: remove the length from the input argument in Sum() and pass a slice to 
Sum() buy calling the funtion like: Sum(arr[:])

The slice-form of the function works for arrays of different lengths!

b) Write a function SumAndAverage which returns these two as unnamed variables
of type int and float32 */
package main

import (
	"fmt"
)

func main() {
	a := [4]float64{2, 3, 4, 5}
	fmt.Println(Sum(a[:]))
	fmt.Println(SumAndAverage(a[:]))
}

func Sum(arrF []float64) float64 {
	sum := float64(0)
	for i := 0; i < len(arrF); i++ {
		sum += arrF[i]
	}
	return sum
}

func SumAndAverage(arrF []float64) (int, float32) {
	sum := float64(0)
	for i := 0; i < len(arrF); i++ {
		sum += arrF[i]
	}
	mySum := int(sum)
	myAvg := float32(sum / float64(len(arrF)))
	return mySum, myAvg
}
