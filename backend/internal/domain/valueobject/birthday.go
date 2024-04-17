package valueobject

import (
	"github.com/take0fit/knowledge-out/pkg/util"
	"time"
)

type Birthday struct {
	Time  time.Time
	Valid bool
}

func NewBirthday(t time.Time) Birthday {
	return Birthday{Time: t}
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
