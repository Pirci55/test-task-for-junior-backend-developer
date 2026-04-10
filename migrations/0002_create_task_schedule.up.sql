CREATE TABLE IF NOT EXISTS task_schedule (
	id BIGSERIAL PRIMARY KEY,
	task_id BIGINT NOT NULL,
    rule_type TEXT NOT NULL,
    start_from DATE NOT NULL DEFAULT CURRENT_DATE,

    interval INT,
    month_day INT,
    is_even BOOLEAN,
    specific_dates DATE[],

	CONSTRAINT fk_tasks_id FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE
);
