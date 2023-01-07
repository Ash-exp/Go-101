package main

import (
	"fmt"
 	"math"
)
func compoundInterest(p float64, t int, r float64, n int){
	a := float64(p) * math.Pow((1+(r/float64(n))), float64(n*t))
	fmt.Printf("The compound interest is : %.2f",a-float64(p))
}

func main() {
	var time,n int
	var principle,rate float64
	fmt.Println("Enter the values of principle, time, rate and n (Space Separated):")
	//Scanln() for line input
	//Scanf() for formated input
	_, err := fmt.Scanf("%f %d %f %d", &principle, &time, &rate, &n)
	if err != nil {
		fmt.Println(err)
	}
	compoundInterest(principle, time, rate, n)
}