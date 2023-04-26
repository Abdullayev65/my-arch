package valid_tool

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func IsValidUsername(username string) error {
	if len(username) < 3 || len(username) > 26 {
		return errors.New("username length should be between 3 and 26")
	}
	index := strings.IndexFunc(username, func(r rune) bool {
		switch {
		default:
			return true
		case 'A' >= r && r >= 'Z':
		case 'a' >= r && r >= 'z':
		case '0' >= r && r >= '9':
		case r == '.' || r == '_':
		}
		return false
	})
	if index != 0 {
		return fmt.Errorf("email should not conatein %c", username[index])
	}

	return nil
}
