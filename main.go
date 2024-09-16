package main

import (
	"fmt"
	"strings"
)

// package level variables:
var (
	remainingTickets uint8 = 30
	bookings               = make([]string, 0, totalTickets) // len=0 (no elements), cap=50
)

const (
	conferenceName       = "Go Conference"
	totalTickets   uint8 = 50
)

func main() {

	// multiple retunrs are allowed in go
	soldTickets := initConference()
	displayBookings()

	greetUsers(soldTickets)

	// infinite loop for ==> for "true", i.e. for { }
	// for remainingTickets > 0 { // loop with true false condition.
	for {
		userFirstName, userLastName, userEmail, userTickets := getUserInputs()
		if vi := isValidInputData(userFirstName, userLastName, userEmail, userTickets); !vi {
			// continue to next iteration to try again:
			continue
		}

		bookTickets(userFirstName, userLastName, userTickets, userEmail, soldTickets)
		soldTickets = calculateSoldTickets(remainingTickets, totalTickets)

		names := getBookingsByPeopleNames(bookings)
		fmt.Println("All Bookings by People names: ", names)

		fmt.Printf("Remaining tickets for %s are %d. Tickets sold %d.\n", conferenceName, remainingTickets, soldTickets)
		if remainingTickets == 0 {
			// end program:
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}
}

func bookTickets(userFirstName string, userLastName string, userTickets uint8, userEmail string, soldTickets uint8) {
	updateBookings(soldTickets, userTickets, userFirstName, userLastName, userEmail)
	displayBookings()

	// recalculate remaining tickets:
	remainingTickets -= userTickets
	fmt.Printf("User %s %s booked %d tickets. Booking details are send to %s\n", userFirstName, userLastName, userTickets, userEmail)
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

func calculateSoldTickets(remainingTickets uint8, totalTickets uint8) uint8 {
	return totalTickets - remainingTickets
}

func initBookings(totalTickets uint8, soldTickets uint8) {
	for i := 0; i < int(totalTickets); i++ {
		if i < int(soldTickets) {
			// The append built-in function appends elements to the end of a slice.
			// If it has sufficient capacity, the destination is resliced to accommodate the new elements.
			// If it does not, a new underlying array will be allocated
			bookings = append(bookings, "SOLD")
		} else {
			// leave empty avaialble indexes so the len can be smaller that the slice capacity.
			break
		}
	}
	// bookings = (len=20, cap=50)
}

func updateBookings(soldTickets uint8, userTickets uint8, userFirstName string, userLastName string, userEmail string) {
	lastBookIndex := soldTickets + userTickets
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
	// Slice Case:
	for i := soldTickets; i < lastBookIndex; i++ {
		bookings = append(bookings, fmt.Sprintf("SOLD-TO %s %s W/E-MAIL %s", userFirstName, userLastName, userEmail))
	}
}

func displayBookings() {
	fmt.Printf("The whole Bookings collection: %v\n", bookings)
	fmt.Print("Bookings display indexed: {")
	for i := 0; i < len(bookings); i++ {
		fmt.Printf("[%d]=%s", i, bookings[i])
		if i == len(bookings)-1 {
			fmt.Println("}")
		} else {
			fmt.Print(", ")
		}
	}
	fmt.Printf("Collection type is %T, length is %d and capacity is %d\n", bookings, len(bookings), cap(bookings))
}

func isValidInputData(userFirstName string, userLastName string, userEmail string, userTickets uint8) (vi bool) {
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

// Booking Slice by reference, i.e. the Slice Descriptor pointing to the underlying array.
func getBookingsByPeopleNames(bookings []string) []string {
	ns := []string{}
	// this is valid as no resize is done in the slice (simply a slice iteration to display values)
	// but if there is a resize/update/new array is built the changes won't be reflected until a
	// new return update the package level variable:
	for _, booking := range bookings {
		// strings.Fields(booking):
		// Fields splits the string s around each instance of one or more consecutive white space characters,
		// as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.
		fields := strings.Fields(booking)
		if len(fields) > 0 && fields[0] != "SOLD" {
			ns = append(ns, fmt.Sprintf("%s %s", fields[1], fields[2]))
		}
	}
	return ns
}
