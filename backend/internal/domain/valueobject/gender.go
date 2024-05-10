package valueobject

import (
	"errors"
)

type Gender struct {
	Value string
	Valid bool
}

func NewGender(value *string) (Gender, error) {
	if value == nil || *value == "" {
		return Gender{Valid: false}, nil
	}
	switch *value {
	case "male", "female", "other":
		return Gender{Value: *value, Valid: true}, nil
	default:
		return Gender{Valid: false}, errors.New("invalid gender value")
	}
}

func (g Gender) String() *string {
	if !g.Valid {
		return nil
	}
	return &g.Value
}
