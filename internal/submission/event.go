package submission

import (
	"context"
	"encoding/json"
	"time"
)

type (
	Broker interface {
		Subscribe(event string, callback func(msg []byte) error)
	}

	EventService interface {
		ReceiveSubmission(ctx context.Context, submission *Submission) error
	}

	Event struct {
		UserID             string    `json:"user_id"`
		QuestionID         string    `json:"question_id"`
		QuestionTitle      string    `json:"question_title"`
		QuestionDifficulty string    `json:"question_difficulty"`
		QuestionLink       string    `json:"question_link"`
		Success            bool      `json:"success"`
		At                 time.Time `json:"at"`
	}

	EventHandler struct {
		event   string
		service EventService
	}
)

func NewEventHandler(event string, service EventService) *EventHandler {
	return &EventHandler{event: event, service: service}
}

func (e *EventHandler) Subscribe(b Broker) {
	b.Subscribe(e.event, e.Read)
}

func (e *EventHandler) Read(msg []byte) error {
	var event Event
	if err := json.Unmarshal(msg, &event); err != nil {
		return err
	}

	return e.service.ReceiveSubmission(context.Background(), &Submission{
		Question: Question{
			ID:         event.QuestionID,
			Title:      event.QuestionTitle,
			Difficulty: event.QuestionDifficulty,
			Link:       event.QuestionLink,
		},
		Success: event.Success,
		At:      event.At,
	})
}
