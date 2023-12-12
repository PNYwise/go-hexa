package status

import "errors"

type Enum string

const (
	ACTIVE   Enum = "active"
	INACTIVE Enum = "inactive"
)

func ValidateStatus(status Enum) error {
	switch status {
	case ACTIVE, INACTIVE:
		return nil // Valid status
	default:
		return errors.New("invalid status")
	}
}
