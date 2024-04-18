package model

import (
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
	"time"
)

// MarshalDate converts a time.Time to a string.
func MarshalDate(t time.Time) graphql.Marshaler {
	dateStr := t.Format("2006-01-02") // ISO 8601 full-date format (only date part).
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(dateStr))
	})
}

// UnmarshalDate parses a string to a time.Time.
func UnmarshalDate(v interface{}) (time.Time, error) {
	str, ok := v.(string)
	if !ok {
		return time.Time{}, errors.New("date must be a string")
	}

	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return time.Time{}, errors.New("date must be in YYYY-MM-DD format")
	}

	return t, nil
}
