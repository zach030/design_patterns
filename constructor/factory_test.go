package constructor

import (
	"testing"
)

func TestNewRestaurant(t *testing.T) {
	NewRestaurant("h").GetFood()
	NewRestaurant("n").GetFood()
}