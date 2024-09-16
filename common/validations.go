package common

import (
	"fmt"
	"strings"
)

func IsValidInputData(userFirstName string, userLastName string, userEmail string, userTickets uint8, remainingTickets uint8, conferenceName string) (vi bool) {
	isValidUserName := len(userFirstName) >= 2 && len(userLastName) >= 2
	vi = true
	if !isValidUserName {
		fmt.Println("Please enter a valid name")
		vi = false
	}

	isValidEmail := len(userEmail) > 3 && strings.Contains(userEmail, "@")
	if !isValidEmail {
		fmt.Println("Please enter a valid email address")
		vi = false
	}

	isValidTicket := userTickets > 0
	if !isValidTicket {
		fmt.Println("Please enter a valid number of tickets")
		vi = false
	}

	// validate avaibaility:
	if userTickets > remainingTickets {
		fmt.Println("Sorry, we only have", remainingTickets, "tickets left for", conferenceName)
		vi = false
	}
	return vi
}
