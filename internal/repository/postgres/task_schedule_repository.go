package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	scheduledomain "example.com/taskservice/internal/domain/schedule"
)

func (r *Repository) CreateSchedule(ctx context.Context, tx any, schedule *scheduledomain.Schedule) (*scheduledomain.Schedule, error) {
	const query = `
		INSERT INTO task_schedule (task_id, rule_type, start_from, interval, month_day, is_even, specific_dates)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, task_id, rule_type, start_from, interval, month_day, is_even, specific_dates
	`

	executor := r.getDB(tx)

	row := executor.QueryRow(
		ctx, query,
		schedule.TaskID,
		schedule.RuleType,
		schedule.StartFrom,
		schedule.Interval,
		schedule.MonthDay,
		schedule.IsEven,
		schedule.SpecificDates,
	)
	created, err := scanSchedule(row)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (r *Repository) ScheduleList(ctx context.Context, tx any, taskId int64) ([]scheduledomain.Schedule, error) {
	const query = `
		SELECT id, task_id, rule_type, start_from, interval, month_day, is_even, specific_dates
		FROM task_schedule
		WHERE task_id = $1
	`

	executor := r.getDB(tx)

	rows, err := executor.Query(ctx, query, taskId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	schedules := make([]scheduledomain.Schedule, 0)
	for rows.Next() {
		schedule, err := scanSchedule(rows)
		if err != nil {
			return nil, err
		}

		schedules = append(schedules, *schedule)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *Repository) GetScheduleByID(ctx context.Context, tx any, scheduleId int64) (*scheduledomain.Schedule, error) {
	const query = `
		SELECT id, task_id, rule_type, start_from, interval, month_day, is_even, specific_dates
		FROM task_schedule
		WHERE id = $1
	`

	executor := r.getDB(tx)

	row := executor.QueryRow(ctx, query, scheduleId)
	found, err := scanSchedule(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, scheduledomain.ErrNotFound
		}

		return nil, err
	}

	return found, nil
}

func (r *Repository) UpdateSchedule(ctx context.Context, tx any, schedule *scheduledomain.Schedule) (*scheduledomain.Schedule, error) {
	const query = `
		UPDATE task_schedule
		SET rule_type = $1,
			start_from = $2,
			interval = $3,
			month_day = $4,
			is_even = $5,
			specific_dates = $6
		WHERE id = $7
		RETURNING id, task_id, rule_type, start_from, interval, month_day, is_even, specific_dates
	`

	executor := r.getDB(tx)

	row := executor.QueryRow(
		ctx, query,
		schedule.RuleType,
		schedule.StartFrom,
		schedule.Interval,
		schedule.MonthDay,
		schedule.IsEven,
		schedule.SpecificDates,
		schedule.ID,
	)
	updated, err := scanSchedule(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, scheduledomain.ErrNotFound
		}

		return nil, err
	}

	return updated, nil
}

func (r *Repository) DeleteSchedule(ctx context.Context, tx any, scheduleId int64) error {
	const query = `DELETE FROM task_schedule WHERE id = $1`

	executor := r.getDB(tx)

	result, err := executor.Exec(ctx, query, scheduleId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return scheduledomain.ErrNotFound
	}

	return nil
}


type scheduleScanner interface {
	Scan(dest ...any) error
}

func scanSchedule(scanner scheduleScanner) (*scheduledomain.Schedule, error) {
	var (
		schedule  scheduledomain.Schedule
		rule_type string
	)

	if err := scanner.Scan(
		&schedule.ID,
		&schedule.TaskID,
		&rule_type,
		&schedule.StartFrom,
		&schedule.Interval,
		&schedule.MonthDay,
		&schedule.IsEven,
		&schedule.SpecificDates,
	); err != nil {
		return nil, err
	}

	schedule.RuleType = scheduledomain.Rule(rule_type)

	return &schedule, nil
}
