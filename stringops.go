package main

import (
	"fmt"
	"strings"
)

func main(){
	var str string = "Hello"

	fmt.Printf("%s\n",str)

	fmt.Println(len(str))  
	con :=  []string{str,"world!"}   
	fmt.Println(strings.Join(con, " "))
}