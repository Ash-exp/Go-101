package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter a number:")
	fmt.Scanln(&n)
	counter := 0
	for t := n;t > 0; {
		counter += 1
		t = t/10	
	}
	fmt.Printf("%d has %d digits",n ,counter)
}