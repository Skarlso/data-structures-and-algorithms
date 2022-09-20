package chapter19

import "strings"

func MakeUppercase(list []string) []string {
	var result []string
	for _, w := range result {
		result = append(result, strings.ToUpper(w))
	}
	return result
}

func MakeUppercaseO1(list []string) {
	for i, w := range list {
		list[i] = strings.ToUpper(w)
	}
}
