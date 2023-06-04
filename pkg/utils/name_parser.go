package utils

import (
	"strings"
)

func ToCamelCase(name string) string {
	var result strings.Builder
	for i := 0; i < len(name)-1; i++ {
		if name[i] >= byte('A') && name[i] <= byte('Z') && i != 0 {
			result.WriteByte('_')
		}
		result.WriteByte(name[i])
	}
	result.WriteByte(name[len(name)-1])
	return strings.ToLower(result.String())
}
