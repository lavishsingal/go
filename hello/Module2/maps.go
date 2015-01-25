package main


import "fmt"

type Vertex struct{
	Lat , Long float64
}

var m map[string]Vertex

func main(){
	m = make(map[string]Vertex)
	m["Map Reduce"] = Vertex{
		42.34 , 43.543,
	}
	fmt.Println(m["Map Reduce"])

}
