package valueobject

import (
	"github.com/take0fit/knowledge-out/pkg/util"
	"time"
)

type Birthday struct {
	Time  time.Time
	Valid bool
}

func NewBirthday(t *string) Birthday {
	if t == nil {
		return Birthday{Valid: false}
	}
	parsedTime, err := time.Parse("2006-01-02", *t)
	if err != nil {
		return Birthday{Valid: false}
	}
	return Birthday{Time: parsedTime, Valid: true}
}

func (b Birthday) Age() *int {
	if !b.Valid {
		return nil
	}

	now := time.Now()
	years := now.Year() - b.Time.Year()

	if util.IsBefore(b.Time, now) {
		years--
	}

	return &years
}

func (b Birthday) String() *string {
	if !b.Valid {
		return nil
	}
	formatted := b.Time.Format("2006-01-02")
	return &formatted
}
