package mergestruct

import (
	"time"

	"github.com/AdrianLungu/decimal"
)

// Event some event
type Event struct {
	ID        int64
	Name      string
	Number    decimal.Decimal
	PNumber   *decimal.Decimal
	Detail    EventDetail
	EmittedAt time.Time
}

// EventDetail details
type EventDetail struct {
	UserID      int64
	Description string
	Tags        []string
}

func f(f func(int) int, i int) int {
	return f(i)
}
