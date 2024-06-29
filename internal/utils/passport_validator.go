package utils

import "strconv"

func ValidatePassword(password string) bool {
	// 1111 111111
	if len(password) != 11 {
		return false
	}
	if password[4] != ' ' {
		return false
	}

	left := password[:4]
	if _, err := strconv.Atoi(left); err != nil {
		return false
	}

	right := password[5:]
	if _, err := strconv.Atoi(right); err != nil {
		return false
	}

	return true
}
