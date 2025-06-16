package logger

import (
	"context"
	"encoding/json"
	"github.com/biryanim/hezzl_tz/internal/model"
	"github.com/biryanim/hezzl_tz/internal/repository"
	"github.com/biryanim/hezzl_tz/internal/service"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
	"time"
)

const (
	batchSize    = 10
	flushTimeout = 5 * time.Second
)

type serv struct {
	nc      *nats.Conn
	repo    repository.LogRepository
	subject string
	buffer  []model.LogEntry
	mu      sync.Mutex
}

func NewService(nc *nats.Conn, repo repository.LogRepository, subject string) service.LoggerService {
	return &serv{
		nc:      nc,
		repo:    repo,
		subject: subject,
		buffer:  make([]model.LogEntry, 0, batchSize),
	}
}

func (s *serv) Run(ctx context.Context) {
	sub, err := s.nc.Subscribe(s.subject, s.handleMessage)
	if err != nil {
		log.Printf("failed to subscribe to subject %s: %v", s.subject, err)
		return
	}
	defer func() {
		if err := sub.Unsubscribe(); err != nil {
			log.Printf("failed to unsubscribe from subject %s: %v", s.subject, err)
		}
	}()

	ticker := time.NewTicker(flushTimeout)
	defer ticker.Stop()

	log.Printf("logger service started, listening on subject %s", s.subject)
	for {
		select {
		case <-ctx.Done():
			log.Println("logger service stopping...")
			s.flush(context.Background())
			return
		case <-ticker.C:
			s.flush(ctx)
		}
	}
}

func (s *serv) handleMessage(msg *nats.Msg) {
	var entry model.LogEntry
	if err := json.Unmarshal(msg.Data, &entry); err != nil {
		log.Printf("failed to unmarshal log message: %v", err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.buffer = append(s.buffer, entry)
	if len(s.buffer) >= batchSize {
		s.flushUnsafe(context.Background())
	}
}

func (s *serv) flush(ctx context.Context) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.flushUnsafe(ctx)
}

func (s *serv) flushUnsafe(ctx context.Context) {
	if len(s.buffer) == 0 {
		return
	}

	log.Printf("flushing %d entries", len(s.buffer))
	if err := s.repo.Add(ctx, s.buffer); err != nil {
		log.Printf("failed to add entries: %v", err)
	}

	s.buffer = make([]model.LogEntry, 0, batchSize)
}
