package main

import "fmt"

func main() {
	var a int
	fmt.Printf("Enter a number: ")
    _, err := fmt.Scanf("%d", &a)
	if err != nil {
		fmt.Println(err)
	}
	flag := 0
	if a < 2 {
		fmt.Println("Not A Prime Number")
		flag = 1
	}
	for i := 2; i*i <= a; i++ {
		if a%i==0 {
			fmt.Println("Not A Prime Number")
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("Prime Number")
	}
}