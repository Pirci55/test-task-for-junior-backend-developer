package task

import (
	"context"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type Repository interface {
	Begin(ctx context.Context) (any, error)
    Commit(ctx context.Context, tx any) error
    Rollback(ctx context.Context, tx any) error

	CreateTask(ctx context.Context, tx any, task *taskdomain.Task) (*taskdomain.Task, error)
	GetTaskByID(ctx context.Context, tx any, id int64) (*taskdomain.Task, error)
	UpdateTask(ctx context.Context, tx any, task *taskdomain.Task) (*taskdomain.Task, error)
	DeleteTask(ctx context.Context, tx any, id int64) error
	TaskList(ctx context.Context, tx any) ([]taskdomain.Task, error)
}

type Usecase interface {
	CreateTask(ctx context.Context, input CreateTaskInput) (*taskdomain.Task, error)
	GetTaskByID(ctx context.Context, id int64) (*taskdomain.Task, error)
	UpdateTask(ctx context.Context, id int64, input UpdateTaskInput) (*taskdomain.Task, error)
	DeleteTask(ctx context.Context, id int64) error
	TaskList(ctx context.Context) ([]taskdomain.Task, error)
}

type CreateTaskInput struct {
	Title       string
	Description string
	Status      taskdomain.Status
}

type UpdateTaskInput struct {
	Title       string
	Description string
	Status      taskdomain.Status
}
