package main 

import "fmt"

func main(){
	fmt.Println("Testtest")

	var start int = 1
	var current int = 11

	fmt.Println("Start is", start, "and current is", current) 
	fmt.Printf("Start is %d and current is %d\n", start, current)

	for i := 1; i < 7; i++ {
		fmt.Printf("The %dth number in teh sequence is %d\n", i, current)

		
	}



}