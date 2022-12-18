package helpers

import (
	"regexp"
)

func VerifyEmailDomain(email string) bool {
	emailregex := regexp.MustCompile(".*@iitk.ac.in$")
	return emailregex.MatchString(email)
}
