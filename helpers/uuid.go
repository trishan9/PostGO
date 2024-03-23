package helpers

import "github.com/google/uuid"

func GenerateOtp() string {
	return uuid.New().String()[:8]
}

func GenerateUid() string {
	return uuid.New().String()
}
