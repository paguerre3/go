package common

import (
	"fmt"

	"github.com/paguerre3/gocomp/domain"
)

// Booking Slice by reference, i.e. the Slice Descriptor pointing to the underlying array.
func GetBookingsByPeopleNames(bookings []domain.BookingUser) []string {
	ns := []string{} // ns is a local variable declared within "the function"
	// this is valid as no resize is done in the slice (simply a slice iteration to display values instead or appending)
	// but if there is a resize/update then a new array is built and changes won't be reflected until because a
	// variable is in the Local Function Frame of the Memory Stack unless a "new return" is providid
	// and then updating the package level variable (otherwise remove passing the slice as a reference
	// so the local/shadowed variable problem is avoided):
	for _, booking := range bookings {
		/*
			// Legacy string case:
			// strings.Fields(booking):
			// Fields splits the string s around each instance of one or more consecutive white space characters,
			// as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.
			fields := strings.Fields(booking) // fields is a local variable declared within the "for" block
			if len(fields) > 0 && fields[0] != "SOLD" {
				ns = append(ns, fmt.Sprintf("%s %s", fields[1], fields[2]))
			}
		*/
		/*
			// Map case:
			if booking["firstName"] != "SOLD" {
				ns = append(ns, fmt.Sprintf("%s %s", booking["firstName"], booking["lastName"]))
			}
		*/
		// Interface & Structs case:
		if booking.FirstName() != "SOLD" {
			ns = append(ns, fmt.Sprintf("%s %s", booking.FirstName(), booking.LastName()))
		}
	}
	return ns
}

func DisplayBookings(bookings []domain.BookingUser) {
	fmt.Printf("The whole Bookings collection: %v\n", bookings)
	fmt.Print("Bookings display indexed: {")
	for i := 0; i < len(bookings); i++ {
		fmt.Printf("[%d]=%v", i, bookings[i])
		if i == len(bookings)-1 {
			fmt.Println("}")
		} else {
			fmt.Print(", ")
		}
	}
	fmt.Printf("Collection type is %T, length is %d and capacity is %d\n", bookings, len(bookings), cap(bookings))
}
