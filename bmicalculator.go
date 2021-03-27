package main;

import "fmt"

func main(){
	var height float32
	var weight float32

	fmt.Println("Enter Weight in Kilograms: ")

	fmt.Scanln(&weight)

	fmt.Println("Enter Height in Meter: ")

	fmt.Scanln(&height)

	bmi := weight/(height*height)
	
	fmt.Print("BMI value: ",bmi)

}