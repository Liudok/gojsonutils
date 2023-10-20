package jsonutils

import (
	"math/rand"
	"strings"
)

const alphanumeric = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const digits = "0123456789"


// RandomString generates a random string of length n (letters only)
func (t* JsonUtils) RandomString(n int) string {
	var sb strings.Builder
	k := len(letters)

	for i := 0; i < n; i++ {
		c := letters[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}


// RandomAlphaNumeric generates a random string of length n of numbers, lower and upper case letters
func (t* JsonUtils) RandomAlphaNumeric(n int) string {
	var sb strings.Builder
	k := len(alphanumeric)

	for i := 0; i < n; i++ {
		c := alphanumeric[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}


// RandomStringOfDigits generates a random string of digits of length n
func (t* JsonUtils) RandomStringOfDigits(n int) string {
	var sb strings.Builder
	k := len(digits)

	for i := 0; i < n; i++ {
		c := digits[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}