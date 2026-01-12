package benchmarking

import (
	"bytes"
	"strings"
)

// CatPlus uses + operator (inefficient for many strings)
func CatPlus(parts []string) string {
	s := ""
	for _, p := range parts {
		s += p
	}
	return s
}

// CatBuilder uses strings.Builder (efficient)
func CatBuilder(parts []string) string {
	var sb strings.Builder
	for _, p := range parts {
		sb.WriteString(p)
	}
	return sb.String()
}

// CatBuffer uses bytes.Buffer (older way, still okay)
func CatBuffer(parts []string) string {
	var buf bytes.Buffer
	for _, p := range parts {
		buf.WriteString(p)
	}
	return buf.String()
}
