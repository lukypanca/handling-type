package model

import "time"

type HandlingSetting struct {
	ID           int
	TipeHandling string
	DescHandling *string
	StartOD      int
	EndOD        int
	Status       string
	IsActive     int
	InsertDate   time.Time
	InsertBy     string
	UpdateDate   time.Time
	UpdateBy     string
}
