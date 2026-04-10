package schedule

import (
	"context"
	"fmt"
	"time"

	scheduledomain "example.com/taskservice/internal/domain/schedule"
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


func (s *Service) CreateSchedule(ctx context.Context, taskId int64, input CreateScheduleInput) (*scheduledomain.Schedule, error) {
	if taskId <= 0 {
		return nil, fmt.Errorf("%w: task id must be positive", ErrInvalidInput)
	}

	normalized, err := validateCreateScheduleInput(input)
	if err != nil {
		return nil, err
	}

	tx, err := s.repo.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer s.repo.Rollback(ctx, tx)

	task, err := s.repo.GetTaskByID(ctx, tx, taskId)
	if err != nil {
		return nil, err
	}

	schedule := &scheduledomain.Schedule{
		TaskID:        task.ID,
		RuleType:      normalized.RuleType,
		StartFrom:     normalized.StartFrom,
		Interval:      normalized.Interval,
		MonthDay:      normalized.MonthDay,
		IsEven:        normalized.IsEven,
		SpecificDates: normalized.SpecificDates,
	}

	createdSchedule, err := s.repo.CreateSchedule(ctx, tx, schedule)
	if err != nil {
		return nil, err
	}

	err = s.repo.Commit(ctx, tx)
	if err != nil {
		return nil, err
	}

	return createdSchedule, nil
}

func (s *Service) GetScheduleByID(ctx context.Context, scheduleId int64) (*scheduledomain.Schedule, error) {
	if scheduleId <= 0 {
		return nil, fmt.Errorf("%w: schedule id must be positive", ErrInvalidInput)
	}

	return s.repo.GetScheduleByID(ctx, nil, scheduleId)
}

func (s *Service) UpdateSchedule(ctx context.Context, scheduleId int64, input UpdateScheduleInput) (*scheduledomain.Schedule, error) {
	if scheduleId <= 0 {
		return nil, fmt.Errorf("%w: schedule id must be positive", ErrInvalidInput)
	}

	normalized, err := validateUpdateScheduleInput(input)
	if err != nil {
		return nil, err
	}

	scheduleModel := &scheduledomain.Schedule{
		ID:            scheduleId,
		RuleType:      normalized.RuleType,
		StartFrom:     normalized.StartFrom,
		Interval:      normalized.Interval,
		MonthDay:      normalized.MonthDay,
		IsEven:        normalized.IsEven,
		SpecificDates: normalized.SpecificDates,
	}

	updatedSchedule, err := s.repo.UpdateSchedule(ctx, nil, scheduleModel)
	if err != nil {
		return nil, err
	}

	return updatedSchedule, nil
}

func (s *Service) DeleteSchedule(ctx context.Context, scheduleId int64) error {
	if scheduleId <= 0 {
		return fmt.Errorf("%w: schedule id must be positive", ErrInvalidInput)
	}

	return s.repo.DeleteSchedule(ctx, nil, scheduleId)
}

func (s *Service) ScheduleList(ctx context.Context, taskId int64) ([]scheduledomain.Schedule, error) {
	if taskId <= 0 {
		return nil, fmt.Errorf("%w: task id must be positive", ErrInvalidInput)
	}

	return s.repo.ScheduleList(ctx, nil, taskId)
}


func validateCreateScheduleInput(input CreateScheduleInput) (CreateScheduleInput, error) {
	if input.Interval == 0 && 
		input.MonthDay == 0 && 
		input.IsEven == nil && 
		len(input.SpecificDates) == 0 {
		return CreateScheduleInput{}, fmt.Errorf("%w: at least one field must be provided", ErrInvalidInput)
	}

	if input.RuleType == "" {
		return CreateScheduleInput{}, fmt.Errorf("%w: rule type is required", ErrInvalidInput)
	}
	if !input.RuleType.Valid() {
		return CreateScheduleInput{}, fmt.Errorf("%w: invalid rule type", ErrInvalidInput)
	}

	if input.Interval < 0 {
		return CreateScheduleInput{}, fmt.Errorf("%w: invalid interval", ErrInvalidInput)
	}
	if input.MonthDay < 0 || input.MonthDay > 30 {
		return CreateScheduleInput{}, fmt.Errorf("%w: invalid month day", ErrInvalidInput)
	}

	return input, nil
}

func validateUpdateScheduleInput(input UpdateScheduleInput) (UpdateScheduleInput, error) {
	if input.Interval == 0 && 
		input.MonthDay == 0 && 
		input.IsEven == nil && 
		len(input.SpecificDates) == 0 {
		return UpdateScheduleInput{}, fmt.Errorf("%w: at least one field must be provided", ErrInvalidInput)
	}

	if input.RuleType == "" {
		return UpdateScheduleInput{}, fmt.Errorf("%w: rule type is required", ErrInvalidInput)
	}
	if !input.RuleType.Valid() {
		return UpdateScheduleInput{}, fmt.Errorf("%w: invalid rule type", ErrInvalidInput)
	}

	if input.Interval < 0 {
		return UpdateScheduleInput{}, fmt.Errorf("%w: invalid interval", ErrInvalidInput)
	}
	if input.MonthDay < 0 || input.MonthDay > 30 {
		return UpdateScheduleInput{}, fmt.Errorf("%w: invalid month day", ErrInvalidInput)
	}

	return input, nil
}
