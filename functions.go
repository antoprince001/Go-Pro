package main

import "fmt"

func main(){
	var a int = 5
	var b int = 6

	sum:= add(a,b)

	fmt.Println("Sum is",sum)
}

func add(a int, b int) int{
	return a+b
}