package model

import "time"

type Showtime struct {
	ID              int64
	MovieID         int
	StudioID        int
	ShowDate        time.Time
	StartTime       time.Time
	Status          string
	DurationMinutes int
}
