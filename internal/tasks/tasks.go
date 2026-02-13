package tasks

import (
	"fmt"
	"slices"
	"time"

	"github.com/Nafine/task-tracker/internal/model"
	"github.com/Nafine/task-tracker/internal/storage"
)

type Service struct {
	tasks model.Tasks
	maxID int
}

func NewService() (*Service, error) {
	data, err := storage.Load()
	if err != nil {
		return nil, fmt.Errorf("could not load tasks: %w", err)
	}

	s := &Service{
		tasks: data,
	}
	s.calculateMaxID()

	return s, nil
}

func (s *Service) calculateMaxID() {
	if len(s.tasks) == 0 {
		s.maxID = 0
		return
	}

	ids := make([]int, 0, len(s.tasks))
	for _, t := range s.tasks {
		ids = append(ids, t.Id)
	}
	s.maxID = slices.Max(ids)
}

func (s *Service) List() model.Tasks {
	return s.tasks
}

func (s *Service) Add(description string) (int, error) {
	s.maxID++
	newTask := model.Task{
		Id:          s.maxID,
		Description: description,
		Status:      model.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.tasks = append(s.tasks, newTask)
	return newTask.Id, s.save()
}

func (s *Service) Update(id int, description string) (bool, error) {
	for i := range s.tasks {
		if s.tasks[i].Id == id {
			s.tasks[i].Description = description
			s.tasks[i].UpdatedAt = time.Now()
			return true, s.save()
		}
	}
	return false, nil
}

func (s *Service) Delete(id int) (bool, error) {
	for i := range s.tasks {
		if s.tasks[i].Id == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return true, s.save()
		}
	}
	return false, nil
}

func (s *Service) Mark(id int, status model.Status) (bool, error) {
	for i := range s.tasks {
		if s.tasks[i].Id == id {
			s.tasks[i].Status = status
			s.tasks[i].UpdatedAt = time.Now()
			return true, s.save()
		}
	}
	return false, nil
}

func (s *Service) save() error {
	if err := storage.Save(s.tasks); err != nil {
		return fmt.Errorf("storage error: %w", err)
	}
	return nil
}
