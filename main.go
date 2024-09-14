package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const totalTickets uint8 = 50
	var remainingTickets uint8 = 30
	soldTickets := calculateSoldTickets(remainingTickets, totalTickets)
	bookings := make([]string, 0, totalTickets)                  // len=0, cap=50
	bookings = initBookings(totalTickets, soldTickets, bookings) // len=20, cap=50

	displayBookings(bookings)

	// Printing the types of variables:
	fmt.Printf("conferenceName is %T, totalTickets is %T, remainingTickets is %T, soldTickets is %T\n", conferenceName, totalTickets,
		remainingTickets, soldTickets)
	// Notice spaces are added for readability by defalt in go fmt:
	fmt.Println("Pointer of soldTickets is", &soldTickets, "and its Variable value", soldTickets)

	fmt.Println("###Welcome to", conferenceName, "booking app.###")
	fmt.Printf("Remaining tickets for %s are %d. Tickets sold %d. Get your tickets now!\n", conferenceName, remainingTickets, soldTickets)

	for {
		var userFirstName, userLastName, userEmail string
		var userTickets uint8
		fmt.Println("Enter your first name: ")
		// be aware that new line is counts as space so avoid placing the full name in one variable separated by space:
		fmt.Scan(&userFirstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&userLastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&userEmail)
		fmt.Println("Enter the ammount of tickets to book: ")
		fmt.Scan(&userTickets)
		fmt.Printf("User %s %s booked %d tickets. Booking details are send to %s\n", userFirstName, userLastName, userTickets, userEmail)
		bookings = updateBookings(soldTickets, userTickets, bookings, userFirstName, userLastName, userEmail)

		displayBookings(bookings)

		// recalculate remaining tickets:
		remainingTickets -= userTickets
		soldTickets = calculateSoldTickets(remainingTickets, totalTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			// strings.Fields(booking):
			// Fields splits the string s around each instance of one or more consecutive white space characters,
			// as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.
			fn := strings.Fields(booking)[0]
			if len(fn) > 0 && fn != "SOLD" {
				firstNames = append(firstNames, fn)
			}
		}
		fmt.Println("All Bookings by People: ", firstNames)
		fmt.Printf("Remaining tickets for %s are %d. Tickets sold %d.\n", conferenceName, remainingTickets, soldTickets)
	}
}

func calculateSoldTickets(remainingTickets uint8, totalTickets uint8) uint8 {
	return totalTickets - remainingTickets
}

// Bookiing Slice by reference, i.e. the Slice Descriptor pointing to the underlying array.
func initBookings(totalTickets uint8, soldTickets uint8, bookings []string) []string {
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
	return bookings
}

// Bookiing Slice by reference, i.e. the Slice Descriptor pointing to the underlying array.
func updateBookings(soldTickets uint8, userTickets uint8, bookings []string, userFirstName string, userLastName string, userEmail string) []string {
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
		bookings = append(bookings, fmt.Sprintf("SOLD TO %s %s w/E-Mail %s", userFirstName, userLastName, userEmail))
	}
	return bookings
}

func displayBookings(bookings []string) {
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
