package main 

import "fmt"


type PaymentMethod interface{
	ProcessPayment(int)
}

type CreditCard struct{
	name string
}

func (c CreditCard) ProcessPayment(amt int){
	fmt.Println("Credit card payment of ",amt)
}
type Paypal struct{
	name string
}

func (p Paypal) ProcessPayment(amt int){
	fmt.Println("Paypal payment of ",amt)
}

type UPI struct{
	name string
}

func (u UPI) ProcessPayment(amt int){
	fmt.Println("UPI payment of ",amt)
}

func makePayment(pay PaymentMethod,amt int){
	pay.ProcessPayment(amt)
}
func main(){
	var pay PaymentMethod
	var n, amoumt int
	fmt.Println("Choose payment method (1: CreditCard, 2: PayPal, 3: UPI): ")
	fmt.Scan(&n)
	fmt.Println("Good choice now enter the amt to be sent: ")
	fmt.Scan(&amoumt)

	switch n{
	case 1: pay = CreditCard{name: "rupay"}
	case 2: pay = Paypal{name: "Paypal"}
	case 3: pay = UPI{name: "Gpay"}
	default: fmt.Println("Not available")
	
	}

	makePayment(pay,amoumt)
}


