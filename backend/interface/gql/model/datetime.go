package model

import (
	"fmt"
	"io"
	"time"
)

type DateTime struct {
	time.Time
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (t *DateTime) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("DateTime must be a string")
	}
	parsed, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return fmt.Errorf("DateTime must be a valid RFC3339 date")
	}
	t.Time = parsed
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (t DateTime) MarshalGQL(w io.Writer) {
	w.Write([]byte(t.Format(`"` + time.RFC3339 + `"`)))
}
