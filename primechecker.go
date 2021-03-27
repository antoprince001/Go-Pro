package main

import (
		"fmt"
)

func main(){
	var a int
	flag := true
	fmt.Println("Enter number: ")
	fmt.Scanln(&a)

	for i:=2;i<int(a/2);i++{
		if(a%i == 0){
			fmt.Println("Not Prime")
			flag = false
			break
		} 

	}
	if(flag==true){
		fmt.Println("Prime")
	}
}