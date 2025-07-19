package passwords

import (
	"math/rand"
	"time"
)

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()-_=+[]{}|;:'\"<>,.?/~"
)

// using above constant variables create a randomized list of password characters, and store all the values in a string called validChars
func ValidChars(useUppercase, useDigits, useSpecialChars bool) string {
	var chars string
	if useUppercase {
		chars += uppercaseLetters
	}
	if useDigits {
		chars += digits
	}
	if useSpecialChars {
		chars += specialChars
	}
	chars += lowercaseLetters // Always include lowercase letters

	// Randomize the order of characters
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	runes := []rune(chars)
	r.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}

// GeneratePassword generates a password of a specified length using the valid characters
func GeneratePassword(length int, useUppercase, useDigits, useSpecialChars bool) string {
	validChars := ValidChars(useUppercase, useDigits, useSpecialChars)
	if length <= 0 || len(validChars) == 0 {
		return ""
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	password := make([]rune, length)
	for i := range password {
		password[i] = rune(validChars[r.Intn(len(validChars))])
	}
	return string(password)
}
