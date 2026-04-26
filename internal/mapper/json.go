package mapper

import (
	"encoding/json"
)

func ToJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}" // fallback aman
	}
	return string(b)
}