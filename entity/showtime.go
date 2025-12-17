package entity

import (
	"time"
)

type Showtime struct {
	ID        int64
	MovieID   int
	StudioID  int
	ShowDate  time.Time
	StartTime string
	Status    string
}
