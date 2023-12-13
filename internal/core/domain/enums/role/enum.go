package role

import "errors"

type Enum string

const (
	ADMIN  Enum = "admin"
	MEMBER Enum = "member"
)

func Validate(role Enum) error {
	switch role {
	case ADMIN, MEMBER:
		return nil // Valid role
	default:
		return errors.New("invalid role")
	}
}
