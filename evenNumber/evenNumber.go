package main

import "fmt"

func main() {
	var a int
	fmt.Printf("Enter a number: ")
    _, err := fmt.Scanf("%d", &a)
	if err != nil {
		fmt.Println(err)
	}
	
	if a%2==0 {
		fmt.Println(a, "Is A Even Number")
	} else {
		fmt.Println(a, "Is A Odd Number")
	}
}