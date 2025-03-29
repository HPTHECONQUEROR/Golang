package main

import (
	"fmt"
	// "time"
)
func run(){
	fmt.Println("I am executing!")
}

func main(){

	go run()
	fmt.Println("Hello motherfucker!")
	// time.Sleep(1*time.Second)
	go run()

}