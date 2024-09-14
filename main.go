package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const totalTickets uint8 = 50
	var remainingTickets uint8 = 30
	soldTickets := calculateSoldTickets(remainingTickets, totalTickets)
	var bookings [totalTickets]string
	initBookings(totalTickets, soldTickets, &bookings)

	// Printing the types of variables:
	fmt.Printf("conferenceName is %T, totalTickets is %T, remainingTickets is %T, soldTickets is %T\n", conferenceName, totalTickets,
		remainingTickets, soldTickets)
	// Notice spaces are added for readability by defalt in go fmt:
	fmt.Println("Pointer of soldTickets is", &soldTickets, "and its Variable value", soldTickets)

	fmt.Println("###Welcome to", conferenceName, "booking app.###")
	fmt.Printf("Remaining tickets for %s are %d. Tickets sold %d. Get your tickets now!\n", conferenceName, remainingTickets, soldTickets)

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
	updateBookings(soldTickets, userTickets, &bookings, userFirstName, userLastName, userEmail)

	displayBookings(bookings)

	// recalculate remaining tickets:
	remainingTickets -= userTickets
	soldTickets = calculateSoldTickets(remainingTickets, totalTickets)
	fmt.Printf("Remaining tickets for %s are %d. Tickets sold %d.\n", conferenceName, remainingTickets, soldTickets)
}

func calculateSoldTickets(remainingTickets uint8, totalTickets uint8) uint8 {
	return totalTickets - remainingTickets
}

func initBookings(totalTickets uint8, soldTickets uint8, bookings *[50]string) {
	for i := 0; i < int(totalTickets); i++ {
		if i < int(soldTickets) {
			// must update the actual array value not the pointer, i.e. doing dereference:
			(*bookings)[i] = "SOLD"
		} else {
			(bookings)[i] = "AVAILABLE"
		}
	}
}

func updateBookings(soldTickets uint8, userTickets uint8, bookings *[50]string, userFirstName string, userLastName string, userEmail string) {
	bookIndex := soldTickets + userTickets
	for i := soldTickets; i < bookIndex; i++ {
		// len returns the length of the array pointer dereferenced:
		if int(i) < len(*bookings) {
			// must update the actual array value not the pointer, i.e. doing dereference:
			(*bookings)[i] = fmt.Sprintf("SOLD TO %s %s w/E-Mail %s", userFirstName, userLastName, userEmail)
		}
	}
}

// passing bookings as copy because it won't be updated in the function, simply used for diosplay:
func displayBookings(bookings [50]string) {
	fmt.Printf("The whole Bookings array: %v\n", bookings)
	fmt.Print("Bookings display indexed: {")
	for i := 0; i < len(bookings); i++ {
		fmt.Printf("[%d]=%s", i, bookings[i])
		if i == len(bookings)-1 {
			fmt.Println("}")
		} else {
			fmt.Print(", ")
		}
	}
	fmt.Printf("Array type is %T and its length is %d\n", bookings, len(bookings))
}
