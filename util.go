package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func UserName(firstName, lastName string) string {
	firstName = strings.ToLower(firstName)
	lastName = strings.ToLower(lastName)
	return fmt.Sprintf("%s-%s", firstName, lastName)
}

func GeneratePassword() string {
	return strconv.Itoa(time.Now().Nanosecond())
}
