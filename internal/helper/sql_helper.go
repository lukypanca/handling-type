package helper

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func BuildInClause(values []string) (string, []interface{}) {
	args := make([]interface{}, len(values))
	placeholders := make([]string, len(values))

	for i, v := range values {
		args[i] = v
		placeholders[i] = fmt.Sprintf(":%d",i+1)
	}

	return strings.Join(placeholders, ","), args
}

func ParseString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func ParseTime(s string) *time.Time {
	if s == "" {
		return nil
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return nil
	}
	return &t
}

func ParseInt(s string) *int {
	if s == "" {
		return nil
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}

	return &v
}

func ParseFloat(s string) *float64 {
	if s == "" {
		return nil
	}

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil
	}

	return &v
}