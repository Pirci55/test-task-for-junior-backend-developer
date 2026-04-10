package task

import (
	"context"
	"fmt"
	"strings"
	"time"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type Service struct {
	repo Repository
	now  func() time.Time
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
		now:  func() time.Time { return time.Now().UTC() },
	}
}


func (s *Service) CreateTask(ctx context.Context, input CreateTaskInput) (*taskdomain.Task, error) {
	normalized, err := validateCreateTaskInput(input)
	if err != nil {
		return nil, err
	}

	model := &taskdomain.Task{
		Title:       normalized.Title,
		Description: normalized.Description,
		Status:      normalized.Status,
	}
	now := s.now()
	model.CreatedAt = now
	model.UpdatedAt = now

	created, err := s.repo.CreateTask(ctx, nil, model)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *Service) GetTaskByID(ctx context.Context, id int64) (*taskdomain.Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: task id must be positive", ErrInvalidInput)
	}

	return s.repo.GetTaskByID(ctx, nil, id)
}

func (s *Service) UpdateTask(ctx context.Context, id int64, input UpdateTaskInput) (*taskdomain.Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: task id must be positive", ErrInvalidInput)
	}

	normalized, err := validateUpdateTaskInput(input)
	if err != nil {
		return nil, err
	}

	model := &taskdomain.Task{
		ID:          id,
		Title:       normalized.Title,
		Description: normalized.Description,
		Status:      normalized.Status,
		UpdatedAt:   s.now(),
	}

	updated, err := s.repo.UpdateTask(ctx, nil, model)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *Service) DeleteTask(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("%w: task id must be positive", ErrInvalidInput)
	}

	return s.repo.DeleteTask(ctx, nil, id)
}

func (s *Service) TaskList(ctx context.Context) ([]taskdomain.Task, error) {
	return s.repo.TaskList(ctx, nil)
}


func validateCreateTaskInput(input CreateTaskInput) (CreateTaskInput, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	if input.Title == "" {
		return CreateTaskInput{}, fmt.Errorf("%w: title is required", ErrInvalidInput)
	}

	if input.Status == "" {
		input.Status = taskdomain.StatusNew
	}

	if !input.Status.Valid() {
		return CreateTaskInput{}, fmt.Errorf("%w: invalid status", ErrInvalidInput)
	}

	return input, nil
}

func validateUpdateTaskInput(input UpdateTaskInput) (UpdateTaskInput, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	if input.Title == "" {
		return UpdateTaskInput{}, fmt.Errorf("%w: title is required", ErrInvalidInput)
	}

	if !input.Status.Valid() {
		return UpdateTaskInput{}, fmt.Errorf("%w: invalid status", ErrInvalidInput)
	}

	return input, nil
}
