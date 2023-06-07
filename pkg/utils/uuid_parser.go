package utils

import "github.com/google/uuid"

func StringToUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}
