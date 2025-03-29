package main

import (
	"fmt"
	"time"
)

/*
func w(id int, b chan int, r chan int) {
	for t := range b {
		fmt.Println("w", id, "doing", t)
		r <- t * 2
	}
}
*/

/*
func w(id int, b chan int, r chan int) {
	for t := range b {
		fmt.Println("w", id, "doing", t)
		time.Sleep(time.Millisecond * 500)
		r <- t * 2
	}
	close(r)
}
*/

func w(id int, b <-chan int, r chan<- int) {
	for t := range b {
		fmt.Println("w", id, "doing", t)
		time.Sleep(time.Millisecond * 500)
		r <- t * 2
	}
}

func main() {

	/*
	b := make(chan int)
	r := make(chan int)

	for i := 0; i < 3; i++ {
		go w(i+1, b, r)
	}

	for j := 1; j <= 5; j++ {
		b <- j
	}

	for k := 1; k <= 5; k++ {
		fmt.Println("ans:", <-r)
	}
	*/

	/*
	b := make(chan int, 10)
	r := make(chan int, 10)

	for i := 0; i < 3; i++ {
		go w(i+1, b, r)
	}

	for j := 1; j <= 5; j++ {
		b <- j
	}

	close(b)

	for k := 1; k <= 5; k++ {
		fmt.Println("done:", <-r)
	}
	*/

	b := make(chan int, 10)
	r := make(chan int, 10)

	for i := 0; i < 3; i++ {
		go w(i+1, b, r) 
	}

	for j := 1; j <= 10; j++ {
		b <- j
	}

	close(b)

	for k := 1; k <= 10; k++ {
		fmt.Println("done:", <-r)
	}
}
