package jsonutils

import (
	"fmt"
    "net/mail"
	"regexp"
)

var (
	onlyLetters = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isCapital = regexp.MustCompile(`^[A-Z\s]+$`).MatchString
	isToken = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isLower = regexp.MustCompile(`^[a-z\s]+$`).MatchString
)


func ValidateStringLength(value string, minLength int, maxLength int) error {
	n := len(value)

	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}

	return nil
}

func ValidMailAddress(address string) error {

    _, err := mail.ParseAddress(address)

    if err != nil {
        return fmt.Errorf("is not a valid email address")
    }

    return nil
}

func ValidateName(value string) error {

	if err := ValidateStringLength(value, 3, 100); err != nil {
		return err
	}

	if !onlyLetters(value) {
		return fmt.Errorf("must contain only letters")
	}

	if !isCapital(value[0:1]) {
		return fmt.Errorf("first letter should be capital")
	}
	return nil
}
