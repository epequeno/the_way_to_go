/* Write a function "Season" which has an input-parameter a month-number
and which returns the name of the season to which this month belongs (disregard)
the day in the month). */
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 12; i++ {
		fmt.Println(Season(i))
	}
}

func Season(monthNum int) (season string) {
	switch {
	case monthNum <= 3:
		return "Spring"
	case 3 < monthNum && monthNum <= 6:
		return "Summer"
	case 6 < monthNum && monthNum <= 9:
		return "Fall"
	default:
		return "Winter"
	}
	return
}
