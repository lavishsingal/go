package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sum , prev := 0 , 1
 	return func() int {
		sum , prev = sum+prev , sum
		return prev
	}
	
	
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}


