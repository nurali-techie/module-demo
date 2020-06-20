package util

import (
	"fmt"
	"strings"
)

func UserName(firstName, lastName string) string {
	firstName = strings.ToLower(firstName)
	lastName = strings.ToLower(lastName)
	return fmt.Sprintf("%s-%s", firstName, lastName)
}
