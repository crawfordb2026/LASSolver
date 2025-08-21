package main

import (
	"fmt"
	"os"
)

func main2() {
	numIter := -1
	fmt.Print("How many iterations of the LookAndSay puzzle do you want to print: ")
	fmt.Scan(&numIter)

	if numIter <= 0 {
		fmt.Println("Invalid number of iterations, please choose a number >= 1")
	}

	fmt.Println("The 1st number in the sequence is 1")
	if numIter == 1 {
		os.Exit(0)
	}
	fmt.Println("The 2nd number in the sequence is 11")
	if numIter == 2 {
		os.Exit(0)
	} // hardcoding 1-3 because they have a different suffix than the rest
	fmt.Println("The 3rd number in the sequence is 21")

	current := make([]int, 20)
	current[0] = 2
	current[1] = 1
	// current[2] = 1
	// current[3] = 1

	next := make([]int, 20)
	var index int 
	for i := 4; i <= numIter; i++ { // this loop covers the NUMBER of iterations of the puzzle
		var numRepeat int = 1
		var prevNum int = current[0]
		index = 0 // current index for next
		
		// fmt.Print("numRepeat: ", numRepeat, "\n")
		// fmt.Print("prevNum: ", prevNum, "\n")

		//printCurrent(current)

		for j := 1; j < len(current); j++ { // this loop iterates thru EACH INDEX of each iteration of the puzzle
			// fmt.Print("Current: ")
			// printCurrent(current)
			// fmt.Print("numRepeat: ", numRepeat, "\n")
			// fmt.Print("prevNum: ", prevNum, "\n")
			if current[j] == prevNum {
				numRepeat++
			} else {
				if index < 20 {
					next[index] = numRepeat
					index++
					next[index] = prevNum
					index++

					prevNum = current[j]
					numRepeat = 1
				}
			}
		} // end of inner loop

		//current = next
			for k := 0; k < len(current); k++ {
				current[k] = next[k] 
			}
		
		fmt.Printf("The %dth number in the sequence is ", i)
		printCurrent(next)

	} // end of outer loop

} // end of main function

func printCurrent2(next []int ){
	for j := 0; j < len(next); j++ {
		if (next[j] == 0){
			break
		}
			
		fmt.Print(next[j])
	}
	fmt.Println()
}
