package goods

import (
	"context"
	"encoding/json"
	"github.com/biryanim/hezzl_tz/internal/model"
	"log"
)

func (s *serv) publishLogEvent(ctx context.Context, good *model.Good) {
	entry := &model.LogEntry{
		Id:          good.ID,
		ProjectId:   good.ProjectID,
		Name:        good.Info.Name,
		Description: good.Info.Description,
		Priority:    good.Priority,
		Removed:     good.Removed,
		EventType:   good.CreatedAt,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		log.Printf("Failed to marshal log entry: %v", err)
		return
	}
	if err = s.nats.Publish(ctx, subject, data); err != nil {
		log.Printf("Failed to publish log entry: %v", err)
		return
	}
}
