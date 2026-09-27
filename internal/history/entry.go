package history

import "time"

type Entry struct {
	Command   string
	Timestamp time.Time
}
