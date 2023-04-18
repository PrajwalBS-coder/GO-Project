package main
import (
	"fmt"


)

func main(){
	var confna="Go Conferenece"
	const tickets int=50
	var remaintickets int =50
	var name string
	var utickets int

fmt.Println("GO Project!")
fmt.Printf("Welcome to %s!\n",confna)
fmt.Printf("Get Your Tickets There are %d Tickets Availabel\n",remaintickets)
fmt.Println("Enter Your Name")
fmt.Scan(&name) 
fmt.Println("Enter the no.of tickets")
fmt.Scan(&utickets)
fmt.Printf("%v booked %v tickets",name,utickets)













}