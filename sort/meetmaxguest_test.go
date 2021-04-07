package sort

import "testing"

func TestMeetMaxGuest(t *testing.T) {
	arrival := []int{900, 600, 700}
	departure := []int{1000, 800, 730}
	expected := 2
	result := MeetMaxGuest(arrival, departure)
	if result != expected {
		t.Errorf("Expected: %v Got: %v", expected, result)
	}
}
