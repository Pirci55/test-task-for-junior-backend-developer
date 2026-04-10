package schedule

import (
	"context"
	"time"

	taskdomain "example.com/taskservice/internal/domain/task"
	scheduledomain "example.com/taskservice/internal/domain/schedule"
)

type Repository interface {
	Begin(ctx context.Context) (any, error)
    Commit(ctx context.Context, tx any) error
    Rollback(ctx context.Context, tx any) error

	GetTaskByID(ctx context.Context, tx any, id int64) (*taskdomain.Task, error)

	CreateSchedule(ctx context.Context, tx any, schedule *scheduledomain.Schedule) (*scheduledomain.Schedule, error)
	GetScheduleByID(ctx context.Context, tx any, scheduleId int64) (*scheduledomain.Schedule, error)
	ScheduleList(ctx context.Context, tx any, taskId int64) ([]scheduledomain.Schedule, error)
	UpdateSchedule(ctx context.Context, tx any, schedule *scheduledomain.Schedule) (*scheduledomain.Schedule, error)
	DeleteSchedule(ctx context.Context, tx any, scheduleId int64) error
}

type Usecase interface {
	CreateSchedule(ctx context.Context, taskId int64, input CreateScheduleInput) (*scheduledomain.Schedule, error)
	GetScheduleByID(ctx context.Context, scheduleId int64) (*scheduledomain.Schedule, error)
	UpdateSchedule(ctx context.Context, scheduleId int64, input UpdateScheduleInput) (*scheduledomain.Schedule, error)
	DeleteSchedule(ctx context.Context, scheduleId int64) error
	ScheduleList(ctx context.Context, taskId int64) ([]scheduledomain.Schedule, error)
}

type CreateScheduleInput struct {
    RuleType      scheduledomain.Rule
    StartFrom     time.Time
    Interval      int
    MonthDay      int
    IsEven        *bool
    SpecificDates []time.Time
}

type UpdateScheduleInput struct {
    RuleType      scheduledomain.Rule
    StartFrom     time.Time
    Interval      int
    MonthDay      int
    IsEven        *bool
    SpecificDates []time.Time
}
