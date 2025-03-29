package main

import (
	"fmt"


)

func genNum(ch chan int){
	for i := 1; i<6;i++{
		ch <- i
	}
	close(ch)
}



func main(){
	ch := make(chan int)
	go genNum(ch)
	
	
	for val:= range ch{
		fmt.Print(val)
	}

	
}