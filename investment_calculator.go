package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount := 1000.00 //declaring multiple variables in one line
	expectedReturnRate := 5.5   //another way of declaring a variable
	years := 10.0

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}
