package audit

import (
	"context"
	"time"
)

type Info struct {
	Now  time.Time
	User string
}

func FromContext(ctx context.Context) Info {
	user, _ := ctx.Value("user").(string)

	if user == "" {
		user = "system"
	}

	return Info{
		Now:  time.Now().UTC(),
		User: user,
	}
}
