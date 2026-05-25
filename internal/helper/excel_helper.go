package helper

import (
	"strings"
	"time"
)

// helper
func GetStr(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
func GetInt(i *int) interface{} {
	if i != nil {
		return *i
	}
	return ""
}
func GetTime(t *time.Time) interface{} {
	if t != nil {
		return t.Format("2006-01-02")
	}
	return ""
}
func GetFloat(f *float64) interface{} {
	if f != nil {
		return *f
	}
	return nil
}
func GetRow(row []string, i int) string {
	if i < len(row) {
		return strings.TrimSpace(row[i])
	}
	return ""
}