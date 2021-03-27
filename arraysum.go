package main

import "fmt"

func main(){

	var arr [4] int
	arr2 := [...] int {1,2,3,4,5}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr3:= append(arr[1:2],4)
	fmt.Println(arr3)
	fmt.Println(arr[0] + arr[1] + arr[2])

	sum:=0
	for i:=0;i<len(arr2);i++{
		sum+=arr2[i]
	}
	fmt.Println(sum)

}