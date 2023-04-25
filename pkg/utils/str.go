package utils

import "strings"

func Concat(v ...string) string {
	var sb strings.Builder
	for _, j := range v {
		sb.WriteString(j)
	}
	return sb.String()
}
