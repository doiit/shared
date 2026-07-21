package helper

import (
	"github.com/gofrs/uuid/v5"
)

// GenerateUUID generates a UUID version 7 based on the current timestamp and random bytes.
func GenerateUUID() uuid.UUID {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	return id
}
