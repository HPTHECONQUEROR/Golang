package main

import (
	"errors"
	"fmt"
)
func withdrawAmt(amt int,balance int)(int, error){
	if balance < amt{
		return 0,errors.New("Insufficient funds!")
	} 
	return balance-amt,nil
}


func main(){
	balance := 1000
	var withdraw int

	fmt.Println("Enter the amput to withdraw: ")
	fmt.Scan(&withdraw)
	
	res,err := withdrawAmt(withdraw,balance)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(withdraw,"amt has been withdrawn annd your balance is ",res)
	}


}