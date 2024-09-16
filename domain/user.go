package domain

type BookingUser interface {
	FirstName() string
	LastName() string
	Email() string
	NumberOfTickets() uint8
	IsValidInput(remainingTickets uint8, conferenceName string) bool
}

type user struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint8
}

func NewBookingUser(firstName string, lastName string, email string, numberOfTickets uint8) BookingUser {
	return &user{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: numberOfTickets,
	}
}

func (u *user) FirstName() string {
	return u.firstName
}
func (u *user) LastName() string {
	return u.lastName
}
func (u *user) Email() string {
	return u.email
}
func (u *user) NumberOfTickets() uint8 {
	return u.numberOfTickets
}
