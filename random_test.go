package jsonutils

import "testing"

const testStringLength = 58

func TestJsonUtils_RandomString(t *testing.T) {
	var testUtils JsonUtils

	s := testUtils.RandomString(testStringLength)
	if len(s) != testStringLength {
		t.Error("Length of randomly generated string is wrong")
	}
}

func TestJsonUtils_RandomStringOfDigits(t *testing.T) {
	var testUtils JsonUtils

	s := testUtils.RandomStringOfDigits(testStringLength)
	if len(s) != testStringLength {
		t.Error("Length of randomly generated string is wrong")
	}
	for _, a := range s {
		if a < '0' || a > '9' {
			t.Error("Randomly generated string of digits has wrong symbol in it")
		}
	}
}