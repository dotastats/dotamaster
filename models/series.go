package models

import "time"

type Series struct {
	Id   int
	URL  string
	Time *time.Time
}
