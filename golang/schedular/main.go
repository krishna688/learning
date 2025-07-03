package main

import (
	"fmt"

	"time"
)

// [[8,10], [9,10], [9,11]] = [[9,11] [13,15]]

// 8.00AM - 5PM

// [8.30-9.30]

// [9-10.30, 2-3]

// 8.00AM + 5pm --> remove which atleast one is busy

// Scheduler defines the interface for a scheduling system that manages events for users.

// It provides functionality to find available time slots and book new events.

type Scheduler interface {

	// FindAvailableSlots identifies time periods where all specified users are available.

	//

	// Parameters:

	//   - startTime: The beginning of the time range to search

	//   - endTime: The end of the time range to search

	//   - slotDuration: The minimum duration required for an available slot

	//   - users: List of user IDs to check availability for

	//

	// Returns:

	//   - []TimeSlot: A slice of available time slots that satisfy the duration requirement

	//   - error: An error if the operation fails

	//

	// The function finds all periods within the specified time range where none of the

	// specified users have existing events and the period is at least as long as slotDuration.

	FindAvailableSlots(startTime, endTime time.Time, slotDuration time.Duration, users []string) ([]TimeSlot, error)

	// BookEvent schedules a new event for the specified users.

	//

	// Parameters:

	//   - eventStartTime: The start time of the event to book

	//   - eventEndTime: The end time of the event to book

	//   - users: List of user IDs to include in the event

	//

	// Returns:

	//   - string: A unique identifier for the newly created event

	//   - error: An error if the booking fails

	//

	//

	// This method creates a new event in the system for the specified users and time range.

	// We allow creating overlapping events.

	BookEvent(eventStartTime, eventEndTime time.Time, users []string) (string, error)
}

type TimeSlot struct {
	Start time.Time

	End time.Time
}

func (ts TimeSlot) String() string {

	timeFormat := "15:04"

	if ts.Start.Day() != ts.End.Day() ||

		ts.Start.Month() != ts.End.Month() ||

		ts.Start.Year() != ts.End.Year() {

		timeFormat = "2006-01-02 15:04"

	}

	return fmt.Sprintf("%s to %s",

		ts.Start.Format(timeFormat),

		ts.End.Format(timeFormat))

}

func main() {

	TestFindAvailableSlots()

}

func NewScheduler() Scheduler {

	// Implementation of the Scheduler interface

	return &AvailabileSlots{}

}

var mapUserSlots = map[string][]TimeSlot{}

type AvailabileSlots struct {
}

func (avs *AvailabileSlots) FindAvailableSlots(startTime, endTime time.Time, slotDuration time.Duration, users []string) ([]TimeSlot, error) {

	slots := [][]TimeSlot{}

	for _, user := range users {

		slots = append(slots, mapUserSlots[user])

	}

	fmt.Printf("slots %v", slots)

	for i, val := range slots {

		startTime := val[i].Start

		endTime := val[i].End

		for j := i + 1; j < len(slots); j++ {

			for _, timeSlot := range slots[j] {

				if (timeSlot.Start.After(startTime) || startTime.Equal(timeSlot.Start)) && timeSlot.Start.Before(endTime) {

					endTime = timeSlot.End

				}

			}

		}

		fmt.Println("startTIme ", startTime, "endtime", endTime)

	}

	return nil, nil

}

func (avs *AvailabileSlots) BookEvent(eventStartTime, eventEndTime time.Time, users []string) (string, error) {

	for _, user := range users {

		mapUserSlots[user] = append(mapUserSlots[user], TimeSlot{eventStartTime, eventEndTime})

	}

	return "", nil

}

func TestFindAvailableSlots() {

	scheduler := NewScheduler()

	_, err := scheduler.BookEvent(

		testDate(2023, 5, 15, 9, 0), // 9:00 AM

		testDate(2023, 5, 15, 10, 30), // 10:30 AM

		[]string{"alice", "bob"},
	)

	if err != nil {

		fmt.Printf("Failed to book event: %v", err)

		return

	}

	_, err = scheduler.BookEvent(

		testDate(2023, 5, 15, 14, 0), // 2:00 PM

		testDate(2023, 5, 15, 15, 0), // 3:00 PM

		[]string{"bob", "charlie"},
	)

	if err != nil {

		fmt.Printf("Failed to book event: %v", err)

		return

	}

	fmt.Println(mapUserSlots)

	slots, err := scheduler.FindAvailableSlots(

		testDate(2023, 5, 15, 8, 0), // 8:00 AM

		testDate(2023, 5, 15, 17, 0), // 5:00 PM

		60*time.Minute,

		[]string{"alice", "bob"},
	)

	if err != nil {

		fmt.Printf("Failed to find available slots: %v", err)

		return

	}

	// 1. 8:00 AM - 9:00 AM

	// 2. 10:30 AM - 2:00 PM

	// 3. 3:00 PM - 5:00 PM

	fmt.Printf("Available slots: %v", slots)

}

func testDate(year, month, day, hour, min int) time.Time {

	return time.Date(year, time.Month(month), day, hour, min, 0, 0, time.UTC)

}
