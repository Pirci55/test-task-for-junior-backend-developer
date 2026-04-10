package schedule

import "time"

type Rule string

const (
	RuleDaily    Rule = "daily"
	RuleMonthly  Rule = "monthly"
	RuleSpecific Rule = "specific"
	RuleParity   Rule = "parity"
)

type Schedule struct {
	ID            int64       `json:"id"`
	TaskID        int64       `json:"task_id"`
	RuleType      Rule        `json:"rule_type"`
	StartFrom     time.Time   `json:"start_from"`
	Interval      int         `json:"interval,omitempty"`
	MonthDay      int         `json:"month_day,omitempty"`
	IsEven        *bool       `json:"is_even,omitempty"`
	SpecificDates []time.Time `json:"specific_dates,omitempty"`
}

func (r Rule) Valid() bool {
	switch r {
	case RuleDaily, RuleMonthly, RuleSpecific, RuleParity:
		return true
	default:
		return false
	}
}
