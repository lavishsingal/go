//mutating maps

package main

import "fmt"

// string int pair
func main(){
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value  : " , m["Answer"] )

	delete(m , "Answer")
	fmt.Println("The value is : " , m["Answer"])

	v , ok := m["Answer"]
	fmt.Println("Value " , v  , "Present?" , ok)

}