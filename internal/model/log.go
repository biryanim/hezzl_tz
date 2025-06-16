package model

import "time"

type LogEntry struct {
	Id          int       `ch:"Id" json:"id"`
	ProjectId   int       `ch:"ProjectId" json:"projectId"`
	Name        string    `ch:"Name" json:"name"`
	Description string    `ch:"Description" json:"description,omitempty"`
	Priority    int       `ch:"Priority" json:"priority"`
	Removed     bool      `ch:"Removed" json:"removed"`
	EventType   time.Time `ch:"EventType" json:"eventType"`
}
