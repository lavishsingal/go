package main

import "fmt"

func swap(x , y string)(string , string){
	return y , x
}
func main(){
	fmt.Printf("Hello , world\n");
	a , b := swap("Hello" , "world")
	fmt.Println(a,b);
}
