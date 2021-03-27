package main

import "fmt"

type student struct{
	name string
	rno  int
	dept string
}
func(e student) display() {
    fmt.Println(e.name)
}


func main() {

	var stdata student
	//assign values to members of stdata
	stdata.name = "John"
	stdata.rno = 20
	stdata.dept = "IT"
	
	stdata2 := student{"Anto", 21, "IT"}
	
	fmt.Println(stdata)
	fmt.Println(stdata2)
	stdata.display()
}