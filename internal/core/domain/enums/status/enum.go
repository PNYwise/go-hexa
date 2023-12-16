package status

import (
	"errors"
)

type Enum string

const (
	ACTIVE   Enum = "active"
	INACTIVE Enum = "inactive"
)

func Validate(status Enum) error {
	switch status {
	case ACTIVE, INACTIVE:
		return nil // Valid status
	default:
		return errors.New("invalid status")
	}
}
