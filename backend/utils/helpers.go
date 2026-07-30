package utils

import (
	"strings"
)

func GetAvatarRromUser(first_name, lase_name string) string {
	return strings.ToUpper(string(first_name[0])) + strings.ToUpper(string(lase_name[0]))
}
func GetFullName(first_name, last_name string) string{
	return first_name + " " + last_name
}
