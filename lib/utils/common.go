package utils

import (
	"strings"
)

func ParseBaseUrl(input string) string {
	return strings.Split(input, "@")[0]
}
