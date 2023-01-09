package hilberthotel

import (
	"fmt"
	"time"
)

type WelcomeKit struct {
	BusNumber       int
	PassengerNumber int
	RoomNumber      int
}

func (e WelcomeKit) String() string {
	return fmt.Sprintf("Bus %v - Passenger %v - Room %v", e.BusNumber, e.PassengerNumber, e.RoomNumber)
}
func NewWelcomeKit(busNumber int, passengerNumber int, roomNumber int) WelcomeKit {
	time.Sleep(10 * time.Microsecond)
	return WelcomeKit{busNumber, passengerNumber, roomNumber}
}
