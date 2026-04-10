package handlers

import (
	"time"

	scheduledomain "example.com/taskservice/internal/domain/schedule"
)

type scheduleMutationDTO struct {
	RuleType      scheduledomain.Rule `json:"rule_type"`
	StartFrom     time.Time           `json:"start_from"`
	Interval      int                 `json:"interval,omitempty"`
	MonthDay      int                 `json:"month_day,omitempty"`
	IsEven        *bool               `json:"is_even,omitempty"`
	SpecificDates []time.Time         `json:"specific_dates,omitempty"`
}

type scheduleDTO struct {
	ID            int64               `json:"id"`
	TaskID        int64               `json:"task_id"`
	RuleType      scheduledomain.Rule `json:"rule_type"`
	StartFrom     time.Time           `json:"start_from"`
	Interval      int                 `json:"interval,omitempty"`
	MonthDay      int                 `json:"month_day,omitempty"`
	IsEven        *bool               `json:"is_even,omitempty"`
	SpecificDates []time.Time         `json:"specific_dates,omitempty"`
}

func newScheduleDTO(schedule *scheduledomain.Schedule) scheduleDTO {
	return scheduleDTO{
		ID:            schedule.ID,
		TaskID:        schedule.TaskID,
		RuleType:      schedule.RuleType,
		StartFrom:     schedule.StartFrom,
		Interval:      schedule.Interval,
		MonthDay:      schedule.MonthDay,
		IsEven:        schedule.IsEven,
		SpecificDates: schedule.SpecificDates,
	}
}