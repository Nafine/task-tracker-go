package model

import (
	"fmt"
	"strings"
	"time"
)

type Status int

const (
	StatusTodo Status = iota
	StatusInProgress
	StatusDone
)

func (s Status) String() string {
	switch s {
	case StatusTodo:
		return "todo"
	case StatusInProgress:
		return "in-progress"
	case StatusDone:
		return "done"
	default:
		return "unknown"
	}
}

func ParseStatus(name string) (Status, bool) {
	switch strings.ToLower(name) {
	case "todo":
		return StatusTodo, true
	case "in-progress":
		return StatusInProgress, true
	case "done":
		return StatusDone, true
	default:
		return 0, false
	}
}

func (s Status) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, s.String())), nil
}

func (s *Status) UnmarshalJSON(data []byte) error {
	name := strings.Trim(string(data), `"`)
	val, ok := ParseStatus(name)
	if !ok {
		return fmt.Errorf("invalid status: %s", name)
	}
	*s = val
	return nil
}

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tasks []Task

func (t Task) String() string {
	return fmt.Sprintf("ID: %d | Status: %s | Description: %s\nCreated: %s",
		t.Id, t.Status, t.Description, t.CreatedAt.Format(time.DateTime))
}
