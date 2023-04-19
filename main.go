package main
import (
	"fmt"


)

func main(){
	var confna="Go Conferenece"
	const tickets int=50
	var remaintickets uint =50
	var name,email string
	var utickets uint
	var bookings [50]string

fmt.Println("GO Project!")
fmt.Printf("Welcome to %s!\n",confna)
fmt.Printf("Get Your Tickets There are %d Tickets Availabel\n",remaintickets)





fmt.Println("Enter Your Name")
fmt.Scan(&name)
fmt.Println("Enter Your Email")
fmt.Scan(&email) 
fmt.Println("Enter the no.of tickets")
fmt.Scan(&utickets)
remaintickets=remaintickets-utickets
fmt.Printf("Thank You %v for booking %v tickets you willl confirmation email to this %v email ",name,utickets,email)
fmt.Printf("\n%v tickets remaining for %v",remaintickets,confna)

bookings[0]=name
fmt.Printf("array values %v",bookings)
fmt.Printf("array values %v",bookings[0])











}