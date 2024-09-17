package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidInput(t *testing.T) {
	tests := []struct {
		name             string
		bu               BookingUser
		remainingTickets uint8
		conferenceName   string
		expected         bool
	}{
		{"TestIsValidInput_WithValidUserData_ReturnsTrue", NewBookingUser("John", "Doe", "jdoe@me.com", 10), 10, "Go Conference", true},
		{"TestIsValidInput_WithInvalidUserFirstName_ReturnsFalse", NewBookingUser("J", "Doe", "jdoe@me.com", 10), 10, "Go Conference", false},
		{"TestIsValidInput_WithInvalidUserLastName_ReturnsFalse", NewBookingUser("John", "D", "jdoe@me.com", 10), 10, "Go Conference", false},
		{"TestIsValidInput_WithInvalidUserEmail_ReturnsFalse", NewBookingUser("John", "Doe", "j", 10), 10, "Go Conference", false},
		{"TestIsValidInput_WithInvalidUserNumberOfTickets_ReturnsFalse", NewBookingUser("John", "Doe", "jdoe@me.com", 10), 5, "Go Conference", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.bu.IsValidInput(tt.remainingTickets, tt.conferenceName))
		})
	}
}
