package main
import "fmt"

//defer the execution of a function call until the function 
//that contains the defer statement completes execution
func secondary() {  
    fmt.Println("Inside the secondary()")
}
func main() {  
    //secondary() will be invoked only after executing the statements of main()
    defer secondary()
    fmt.Println("Inside the main()")
}