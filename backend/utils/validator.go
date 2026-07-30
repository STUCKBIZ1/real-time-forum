package utils

import (
	"errors"
	"net/mail"
	"strings"
	"unicode"
)

func IsValidEmail(email string) bool {
	email = strings.TrimSpace(email)
	_, err := mail.ParseAddress(email)
	return err == nil && len(email) < 60
}
func IsValidUserName(username string) bool {
	if len(username) < 3 || len(username) > 20 {
		return false
	}
	runes := []rune(username)
	if !unicode.IsLetter(runes[0]) {
		return false
	}
	for _, r := range runes {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-') {
			return false
		}
	}
	return true
}
func IsValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 72 {
		return false
	}

	var hasUpper, hasLower, hasNumber, hasSymbol bool

	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasNumber = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}
func IsValidFirstLastName(name string) error {
	name = strings.TrimSpace(name)
	runes := []rune(name)

	if len(runes) < 2 || len(runes) > 40 {
		return errors.New("first or last name must be between 2 and 40 characters")
	}
	if !unicode.IsLetter(runes[0]) || !unicode.IsLetter(runes[len(runes)-1]) {
		return errors.New("invalid first or last name")
	}
	for _, r := range runes {
		if !unicode.IsLetter(r) && r != ' ' && r != '-' && r != '\'' {
			return errors.New("invalid first or last name")
		}
	}

	return nil
}
func IsValidAge(age int) error {
	if age < 13 || age > 120 {
		return errors.New("age must be between 13 and 120")
	}
	return nil
}
