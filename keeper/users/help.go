package users

import "strings"

func parseItems(s string) []string {
	const sep = ","
	src := strings.Split(s, sep)
	dst := make([]string, 0)
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		dst = append(dst, item)
	}
	return dst
}

func itemsToString(items []string) string {
	builder := strings.Builder{}
	sep := ""
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		builder.WriteString(sep)
		builder.WriteString(item)
		sep = ","
	}
	return builder.String()
}
