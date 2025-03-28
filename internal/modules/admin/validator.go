package admin

import (
	"errors"
	"regexp"
)

// ValidateAdmin validates the admin data
func ValidateAdmin(admin Admin) error {
	if admin.FirstName == "" {
		return errors.New("firstName is required")
	}
	if admin.LastName == "" {
		return errors.New("lastName is required")
	}
	if admin.EmailID == "" {
		return errors.New("emailId is required")
	}
	if !isValidEmail(admin.EmailID) {
		return errors.New("invalid email format")
	}
	if admin.Password == "" {
		return errors.New("password is required")
	}
	// Add more validation as needed (e.g., mobile format, gender values)
	return nil
}

// isValidEmail checks if the email format is valid
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
