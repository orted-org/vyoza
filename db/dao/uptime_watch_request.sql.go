package db

import (
	"context"
	"time"
)

const addUptimeWatchRequest = `
INSERT INTO uptime_watch_request (
	title,
	description,
	location,
	enabled,
	enable_updated_at,
	interval,
	expected_status,
	max_response_time,
	retain_duration,
	hook_level,
	hook_addr,
	hook_secret
)
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING   id,
			title,
			description,
			location,
			enabled,
			enable_updated_at,
			interval,
			expected_status,
			max_response_time,
			retain_duration,
			hook_level,
			hook_addr,
			hook_secret
`

type AddUptimeWatchRequestParams struct {
	Title           string
	Description     string
	Location        string
	Enabled         bool
	EnableUpdatedAt time.Time
	Interval        int
	ExpectedStatus  int
	MaxResponseTime int
	RetainDuration  string
	HookLevel       int
	HookAddress     string
	HookSecret      string
}

func (q *Queries) AddUptimeWatchRequest(ctx context.Context, arg AddUptimeWatchRequestParams) (UptimeWatchRequest, error) {
	row := q.queryRow(ctx, q.addUptimeWatchRequest, addUptimeWatchRequest, arg.Title, arg.Description, arg.Location, arg.Enabled, time.Now(), arg.Interval, arg.ExpectedStatus, arg.MaxResponseTime, arg.RetainDuration, arg.HookLevel, arg.HookAddress, arg.HookSecret)
	var i UptimeWatchRequest
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Location,
		&i.Enabled,
		&i.EnableUpdatedAt,
		&i.Interval,
		&i.ExpectedStatus,
		&i.MaxResponseTime,
		&i.RetainDuration,
		&i.HookLevel,
		&i.HookAddress,
		&i.HookSecret,
	)
	return i, err
}

const getUptimeWatchRequestByID = `
SELECT id,
    title,
    description,
    location,
    enabled,
    enable_updated_at,
    interval,
    expected_status,
    max_response_time,
    retain_duration,
    hook_level,
    hook_addr,
    hook_secret
FROM uptime_watch_request
WHERE id = ?
`

func (q *Queries) GetUptimeWatchRequestByID(ctx context.Context, id int) (UptimeWatchRequest, error) {
	row := q.queryRow(ctx, q.getUptimeWatchRequestByID, getUptimeWatchRequestByID, id)
	var i UptimeWatchRequest
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Location,
		&i.Enabled,
		&i.EnableUpdatedAt,
		&i.Interval,
		&i.ExpectedStatus,
		&i.MaxResponseTime,
		&i.RetainDuration,
		&i.HookLevel,
		&i.HookAddress,
		&i.HookSecret,
	)
	return i, err
}
