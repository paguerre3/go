package main

import (
	// GO "internal" module:
	"fmt"
	"sync"
	"time"

	// Own module/GO external package:
	"github.com/paguerre3/gocomp/common"
	"github.com/paguerre3/gocomp/domain"
)

// package level variables:
var (
	remainingTickets uint8 = 30
	bookings               = make([]domain.BookingUser, 0, totalTickets) // len=0 (initial size), cap=50
	wg               sync.WaitGroup
)

const (
	conferenceName       = "Go Conference"
	totalTickets   uint8 = 50
)

func main() {
	soldTickets := initConference()
	common.DisplayBookings(bookings)

	greetUsers(soldTickets)

	// infinite loop for ==> for "true", i.e. for { }
	// for remainingTickets > 0 { // loop with true false condition.
	for {
		// multiple returns are allowed in go
		//userFirstName, userLastName, userEmail, userTickets := getUserInputs()
		bookingUser := domain.NewBookingUser(getUserInputs())
		if vi := bookingUser.IsValidInput(remainingTickets, conferenceName); !vi {
			// continue to next iteration to try again:
			continue
		}

		bookTickets(bookingUser, soldTickets)
		wg.Add(1)
		go sendTicket(bookingUser)
		soldTickets = calculateSoldTickets(remainingTickets, totalTickets)

		names := common.GetBookingsByPeopleNames(bookings)
		fmt.Println("All Bookings by People names: ", names)

		fmt.Printf("Remaining tickets for %s are %d. Tickets sold %d.\n", conferenceName, remainingTickets, soldTickets)
		if remainingTickets == 0 {
			// end program:
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}
	fmt.Println("Waiting until all e-mails are sent...")
	wg.Wait()
	fmt.Println("Bye!")
}

func bookTickets(bu domain.BookingUser, soldTickets uint8) {
	updateBookings(soldTickets, bu)
	common.DisplayBookings(bookings)

	// recalculate remaining tickets:
	remainingTickets -= bu.NumberOfTickets()
	fmt.Printf("User %s %s booked %d tickets. Booking details are send to %s\n", bu.FirstName(), bu.LastName(), bu.NumberOfTickets(), bu.Email())
}

func initConference() uint8 {
	soldTickets := calculateSoldTickets(remainingTickets, totalTickets)
	// bookings := make([]string, 0, totalTickets)                  // len=0, cap=50
	// bookings = initBookings(totalTickets, soldTickets, bookings) // len=20, cap=50
	initBookings(totalTickets, soldTickets)
	return soldTickets
}

func greetUsers(soldTickets uint8) {
	// Printing the types of variables:
	fmt.Printf("conferenceName is %T, totalTickets is %T, remainingTickets is %T, soldTickets is %T\n", conferenceName, totalTickets,
		remainingTickets, soldTickets)

	// Notice spaces are added for readability by defalt in go fmt:
	fmt.Println("Pointer of soldTickets is", &soldTickets, "and its Variable value", soldTickets)

	fmt.Println("###Welcome to", conferenceName, "booking app.###")
	fmt.Printf("Remaining tickets for %s are %d. Tickets sold %d. Get your tickets now!\n", conferenceName, remainingTickets, soldTickets)
}

// package level function starts with lower case letter:
func calculateSoldTickets(remainingTickets uint8, totalTickets uint8) uint8 {
	return totalTickets - remainingTickets
}

func initBookings(totalTickets uint8, soldTickets uint8) {
	for i := 0; i < int(totalTickets); i++ {
		if i < int(soldTickets) {
			// The append built-in function appends elements to the end of a slice.
			// If it has sufficient capacity, the destination is resliced to accommodate the new elements.
			// If it does not, a new underlying array will be allocated
			/*
				// legacy string case
				//bookings = append(bookings, "SOLD")
			*/
			/*
				// Map case:
				sd := map[string]string{
					"firstName":       "SOLD",
					"lastName":        "SOLD",
					"email":           "sold@sold.com",
					"numberOfTickets": fmt.Sprintf("%d", soldTickets),
				}
				bookings = append(bookings, sd)
			*/
			// Interface & Structs case:
			sd := domain.NewBookingUser("SOLD", "SOLD", "sold@sold.com", soldTickets)
			bookings = append(bookings, sd)
		} else {
			// leave empty avaialble indexes so the len can be smaller that the slice capacity.
			break
		}
	}
	// bookings = (len=20, cap=50)
}

func updateBookings(soldTickets uint8, bu domain.BookingUser) {
	lastBookIndex := soldTickets + bu.NumberOfTickets()
	/*
		// Array case:
		for i := soldTickets; i < lastBookIndex; i++ {
			// len returns the length of the array pointer dereferenced, i.e. the actual array variable value:
			if int(i) < len(*bookings) {
				// must update the actual array value not the pointer, i.e. Go automatically is doing dereference (*bookings)[i] in case of array:
				bookings[i] = fmt.Sprintf("SOLD TO %s %s w/E-Mail %s", userFirstName, userLastName, userEmail)
			}
		}
	*/
	/*
		// Slice Case:
		for i := soldTickets; i < lastBookIndex; i++ {
			bookings = append(bookings, fmt.Sprintf("SOLD-TO %s %s W/E-MAIL %s", userFirstName, userLastName, userEmail))
		}
	*/
	/*
		// Slice Case using Maps:
		var userData = make(map[string]string)
		userData["firstName"] = userFirstName
		userData["lastName"] = userLastName
		userData["email"] = userEmail
		userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // 10 is used for decimal formatting
		for i := soldTickets; i < lastBookIndex; i++ {
			bookings = append(bookings, userData)
		}
	*/
	// Slice Case using Intreface & Struct:
	/*
		user := domain.User{
			FirstName:       userFirstName,
			LastName:        userLastName,
			Email:           userEmail,
			NumberOfTickets: userTickets,
		}
	*/
	for i := soldTickets; i < lastBookIndex; i++ {
		bookings = append(bookings, bu)
	}
}

func getUserInputs() (string, string, string, uint8) {
	var userFirstName, userLastName, userEmail string
	var userTickets uint8
	fmt.Println("Enter your first name: ")
	// Be aware that new line is counted as space so avoid placing the full name in one variable separated by space, i.e.
	// Scan scans text read from standard input, storing successive space-separated values into successive arguments.
	// Newlines count as space:
	fmt.Scan(&userFirstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&userLastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&userEmail)
	fmt.Println("Enter the ammount of tickets to book: ")
	fmt.Scan(&userTickets)
	return userFirstName, userLastName, userEmail, userTickets
}

func sendTicket(bu domain.BookingUser) {
	defer wg.Done() // defer executes at the end of teh function always
	ticket := fmt.Sprintf("%d tickets for %s %s", bu.NumberOfTickets(), bu.FirstName(), bu.LastName())
	fmt.Printf(">>>Sending ticket:\n %s \nto email address %s\n", ticket, bu.Email())
	// simulates time of sending an email:
	time.Sleep(30 * time.Second)
	fmt.Printf("<<<Ticket sent to email %s!\n", bu.Email())
}
