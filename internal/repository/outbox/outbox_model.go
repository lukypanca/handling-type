package outbox

import "time"

type EventStatus string

const (
	StatusNew       EventStatus = "NEW"
	StatusProcessed EventStatus = "PROCESSED"
	StatusFailed    EventStatus = "FAILED"
)

type OutboxEvent struct {
	ID        int
	EventType string
	Payload   string // JSON string
	Status    EventStatus
	Retry     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
