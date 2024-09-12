package main

import "fmt"

func calculateSoldTickets(remainingTickets uint8, totalTickets uint8) uint8 {
	return totalTickets - remainingTickets
}

func main() {
	var conferenceName = "Go Conference"
	const totalTickets uint8 = 50
	var remainingTickets uint8 = 30
	soldTickets := calculateSoldTickets(remainingTickets, totalTickets)

	// Notice spaces are added for readability by defalt in go fmt:
	fmt.Println("Welcome to", conferenceName, "booking app.")
	fmt.Printf("Remaining tickets for %s are %d. Total tickets sold %d. Get your tickets now!\n", conferenceName, remainingTickets, soldTickets)
}
