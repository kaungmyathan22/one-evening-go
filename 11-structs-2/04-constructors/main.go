package main

import (
	"errors"
	"fmt"
	"time"
)

type DateRange struct {
	Start time.Time
	End   time.Time
}

func (d DateRange) Hours() float64 {
	return d.End.Sub(d.Start).Hours()
}

func NewDateRange(Start, End time.Time) (*DateRange, error) {
	if time.Time.IsZero(Start) || time.Time.IsZero(End) {
		return nil, errors.New("The provided start date or end date is empty")
	}
	if End.Before(Start) {
		return nil, errors.New("The provided end date happens before the start date")
	}
	return &DateRange{
		Start: Start,
		End:   End,
	}, nil
}

func main() {
	lifetime, err := NewDateRange(
		time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
		time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
	)
	if err != nil {
		fmt.Println(lifetime.Hours())
	}

	travelInTime, err := NewDateRange(
		time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
		time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
	)
	if err != nil {
		fmt.Println(travelInTime.Hours())
	}
}
