package domain

import (
	"fmt"
	"strings"
)

func (u *user) IsValidInput(remainingTickets uint8, conferenceName string) (vi bool) {
	isValidUserName := len(u.firstName) >= 2 && len(u.lastName) >= 2
	vi = true
	if !isValidUserName {
		fmt.Println("Please enter a valid name")
		vi = false
	}

	isValidEmail := len(u.email) > 3 && strings.Contains(u.email, "@")
	if !isValidEmail {
		fmt.Println("Please enter a valid email address")
		vi = false
	}

	isValidTicket := u.numberOfTickets > 0
	if !isValidTicket {
		fmt.Println("Please enter a valid number of tickets")
		vi = false
	}

	// validate avaibaility:
	if u.numberOfTickets > remainingTickets {
		fmt.Println("Sorry, we only have", remainingTickets, "tickets left for", conferenceName)
		vi = false
	}
	return vi
}
