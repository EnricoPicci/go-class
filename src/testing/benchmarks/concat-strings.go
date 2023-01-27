package benchmarkspkg

import (
	"fmt"
	"strings"
)

func ConcatWithPlus(strs []string) string {
	var concatenated string

	for _, s := range strs {
		concatenated += s
	}
	return concatenated
}

func ConcatWithStringBuilder(strs []string) string {
	var sb strings.Builder

	for _, s := range strs {
		sb.WriteString(s)
	}
	return sb.String()
}

func ConcatWithStringBuilderAsWriter(strs []string) string {
	var sb strings.Builder

	for _, s := range strs {
		fmt.Fprint(&sb, s)
	}
	return sb.String()
}
