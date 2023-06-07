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

func CommentToString(str string) string {
	cutset := strings.Trim(strings.Split(str, "=")[1], "+")
	return strings.Replace(cutset, "+", " ", -1)
}
