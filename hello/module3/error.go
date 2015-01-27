package main

import (
	"fmt"
	"time"
)

type Myerror struct {
	When time.Time
	Why string 
}

func (e *Myerror ) Error() string{
	return fmt.Sprintf("when %v and why %s" ,e.When , e.Why)
}

func run() error{
	return &Myerror{
		time.Now(),
		"Dint worked",
	}
}

func main(){

	if err := run(); err != nil{
		fmt.Println(err)
	}

}