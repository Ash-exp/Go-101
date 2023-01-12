package main

import "fmt"

func main() {
	var arr [4]int
	fmt.Println("Enter 4 integers :")
	for i := 0; i < 4; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanln(&arr[i])
	}
	large1 := arr[0]
	large2 := arr[0]
	for i := 0; i < 4; i++ {
		if large1 == large2 && arr[i] < large2 {
			large2 = arr[i]
		}
		if large1 < arr[i] {
			large2 = large1
			large1 = arr[i]
		} else if large2 < arr[i] && large1 > arr[i]{
			large2 = arr[i]
		}
	}
	fmt.Printf("The second largest number is : %d",large2)
}