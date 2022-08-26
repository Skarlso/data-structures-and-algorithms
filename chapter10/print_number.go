package chapter10

import "fmt"

func PrintNumber(slice []any) {
	for _, item := range slice {
		if s, ok := item.([]any); ok {
			PrintNumber(s)
		} else {
			fmt.Println(item)
		}
	}
}
