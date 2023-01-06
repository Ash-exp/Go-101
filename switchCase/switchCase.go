package main

import "fmt"

func evenOdd(){
	var a,b int
	fmt.Printf("Enter two numbers: ")
    _, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		fmt.Println(err)
	}
	
	if a%2==0 {
		fmt.Println(a, "Is A Even Number")
	} else {
		fmt.Println(a, "Is A Odd Number")
	}

	if b%2==0 {
		fmt.Println(b, "Is A Even Number")
	} else {
		fmt.Println(b, "Is A Odd Number")
	}
}

func switchEvenOdd(){
	var a,b int
	fmt.Printf("Enter two numbers: ")
    _, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		fmt.Println(err)
	}
	switch a%2==0 {
		case true: 
			fmt.Println(a, "Is An Even Number")
		case false: 
			fmt.Println(a, "Is An Odd Number")
		default: 
			fmt.Println(a, ": Invalid input!")
	}

	switch b%2==0 {
		case true: 
			fmt.Println(b, "Is An Even Number")
		case false: 
			fmt.Println(b, "Is An Odd Number")
		default: 
			fmt.Println(b, ": Invalid input!")
	}
}

func dayName(){
	var day int
	fmt.Printf("Enter a day numbers: ")
    _, err := fmt.Scanf("%d", &day)
	if err != nil {
		fmt.Println(err)
	}
	switch day {
	case 1: 
		fmt.Println("Monday")
	case 2: 
		fmt.Println("Tuesday")
	case 3: 
		fmt.Println("Wednesday")
	case 4: 
		fmt.Println("Thursday")
	case 5: 
		fmt.Println("Friday")
	case 6: 
		fmt.Println("Saturday")
	case 7: 
		fmt.Println("Sunday")
	default: 
		fmt.Println("Invalid input! Please enter week number between 1-7.")
	}
}

func main() {
	evenOdd()
	dayName()
	switchEvenOdd()
}