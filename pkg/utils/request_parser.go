package utils

import (
	"strings"
)

func ArticleToSlice(request string) (result []string) {
	for _, kv := range strings.Split(request, "&") {
		cutset := strings.Trim(strings.Split(kv, "=")[1], "+")
		result = append(result, strings.Replace(cutset, "+", " ", -1))
	}
	return
}
