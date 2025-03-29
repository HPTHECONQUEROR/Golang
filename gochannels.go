// go channels are like a pipeline we can send and recive the data inside it 
// but we have limitations if the channel is unbuffered then we can only pass one go routine inside it at a time 
// if its bufferd then its can allow multiple data flow into it 
// we can also use multiple channels to handle multiple data flow 


// package main
// import (
// 	"fmt"
// 	// "time"
// )

// func putData(ch chan int){
// 	fmt.Println("Putting data into channel..")
// 	ch <- 60
// 	fmt.Println("Job done")
// }

// func main(){

// 	ch := make(chan int)
// 	fmt.Println("Initiating the process")
// 	go putData(ch)
// 	got := <-ch
// 	fmt.Println("Recieevd the data fromt he channel ch: ",got)

// }

package main
import (
	"fmt"
	"time"
)

func putData(ch chan int){
	fmt.Println("Putting data into channel..")
	for i:=0;i<10;i++{
		ch <- i+60
		time.Sleep(time.Millisecond * 200)
	}
	close(ch)
	
	fmt.Println("Job done")
}

func main(){

	ch := make(chan int, 5)
	fmt.Println("Initiating the process")
	go putData(ch)
	for i:=0;i<10;i++{
		got := <-ch
		fmt.Println("Recieevd the data fromt he channel ch: ",got)
	}

	fmt.Println("All data recieved!")
	
	

}