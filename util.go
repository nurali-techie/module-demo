package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func UserName(firstName, lastName, middeName string) string {
	firstName = strings.ToLower(firstName)
	lastName = strings.ToLower(lastName)
	middeName = strings.ToLower(middeName)

	initial := string(middeName[0])
	return fmt.Sprintf("%s-%s-%s", firstName, lastName, initial)
}

func GeneratePassword() string {
	return strconv.Itoa(time.Now().Nanosecond())
}

func GenerateEmail(userName string) string {
	userName = strings.ReplaceAll(userName, "-", ".")
	return fmt.Sprintf("%s@mycompany.com", userName)
}
