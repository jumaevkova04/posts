package helpers

import (
	"errors"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v3"
	"time"
	"unicode"
)

var (
	TimeRFC3339Format = time.RFC3339
)

func NewUUID() string {
	return uuid.New().String()
}

func TimeNowWithRFC3339Format() string {
	return time.Now().Format(TimeRFC3339Format)
}

func NullString(s string) null.String {
	return null.NewString(s, s != "")
}

func StrongPassword(value interface{}) error {
	switch input := value.(type) {
	case string:
		if len(input) > 64 {
			return errors.New("invalid length")
		}

		var (
			hasMinLen  = false
			hasUpper   = false
			hasLower   = false
			hasNumber  = false
			hasSpecial = false
		)
		if len(input) >= 8 {
			hasMinLen = true
		}

		for _, char := range input {
			switch {
			case unicode.IsUpper(char):
				hasUpper = true
			case unicode.IsLower(char):
				hasLower = true
			case unicode.IsNumber(char):
				hasNumber = true
			case unicode.IsPunct(char) || unicode.IsSymbol(char):
				hasSpecial = true
			}
		}

		if hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial {
			return nil
		}
		return errors.New("password is not strong enough")

	default:
		return errors.New("password must be string")
	}
}
