package main

import "fmt"

func main(){
	fmt.Println("Enter number : ")
	var x int
	fmt.Scanln(&x)

	if x%2 == 0{
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}