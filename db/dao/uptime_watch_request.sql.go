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
	ssl_monitor,
	expected_status,
	std_response_time,
	max_response_time,
	retain_duration,
	hook_level,
	hook_addr,
	hook_secret,
	notificationEmail,
)
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	RETURNING id,
	title,
	description,
	location,
	enabled,
	enable_updated_at,
	interval,
	ssl_monitor,
	expected_status,
	std_response_time,
	max_response_time,
	retain_duration,
	hook_level,
	hook_addr,
	hook_secret,
	notificationEmail
`

type AddUptimeWatchRequestParams struct {
	Title             string
	Description       string
	Location          string
	Enabled           bool
	Interval          int
	SSLMonitor        bool
	ExpectedStatus    int
	StdResponseTime   int
	MaxResponseTime   int
	RetainDuration    int
	HookLevel         int
	HookAddress       string
	HookSecret        string
	NotificationEmail string
}

func (q *Queries) AddUptimeWatchRequest(ctx context.Context, arg AddUptimeWatchRequestParams) (UptimeWatchRequest, error) {
	row := q.queryRow(ctx, q.addUptimeWatchRequest, addUptimeWatchRequest, arg.Title, arg.Description, arg.Location, arg.Enabled, time.Now().UTC(), arg.Interval, arg.SSLMonitor, arg.ExpectedStatus, arg.StdResponseTime, arg.MaxResponseTime, arg.RetainDuration, arg.HookLevel, arg.HookAddress, arg.HookSecret)
	var i UptimeWatchRequest
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Location,
		&i.Enabled,
		&i.EnableUpdatedAt,
		&i.Interval,
		&i.SSLMonitor,
		&i.ExpectedStatus,
		&i.StdResponseTime,
		&i.MaxResponseTime,
		&i.RetainDuration,
		&i.HookLevel,
		&i.HookAddress,
		&i.HookSecret,
		&i.NotificationEmail,
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
    std_response_time,
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
		&i.StdResponseTime,
		&i.MaxResponseTime,
		&i.RetainDuration,
		&i.HookLevel,
		&i.HookAddress,
		&i.HookSecret,
	)
	return i, err
}

const getAllUptimeWatchRequest = `
SELECT id,
    title,
    description,
    location,
    enabled,
    enable_updated_at,
    interval,
    expected_status,
    std_response_time,
    max_response_time,
    retain_duration,
    hook_level,
    hook_addr,
    hook_secret
FROM uptime_watch_request
`

func (q *Queries) GetAllUptimeWatchRequest(ctx context.Context) ([]UptimeWatchRequest, error) {
	rows, err := q.query(ctx, q.getAllUptimeWatchRequest, getAllUptimeWatchRequest)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UptimeWatchRequest
	for rows.Next() {
		var i UptimeWatchRequest
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Location,
			&i.Enabled,
			&i.EnableUpdatedAt,
			&i.Interval,
			&i.ExpectedStatus,
			&i.StdResponseTime,
			&i.MaxResponseTime,
			&i.RetainDuration,
			&i.HookLevel,
			&i.HookAddress,
			&i.HookSecret,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteUptimeWatchRequestById = `
DELETE FROM uptime_watch_request
WHERE id = ?
`

func (q *Queries) DeleteUptimeWatchRequestById(ctx context.Context, id int) error {
	_, err := q.exec(ctx, q.deleteUptimeWatchRequestById, deleteUptimeWatchRequestById, id)
	return err
}

func (q *Queries) UpdateUptimeWatchRequestById(ctx context.Context, updateData map[string]interface{}, id int) (UptimeWatchRequest, error) {
	var i UptimeWatchRequest
	updateData["enable_updated_at"] = "?"
	closing := `
	WHERE id = ?
	RETURNING   
		id,
		title,
		description,
		location,
		enabled,
		enable_updated_at,
		interval,
		expected_status,
		std_response_time,
		max_response_time,
		retain_duration,
		hook_level,
		hook_addr
		hook_secret
	`
	qry, err := CreateDynamicUpdateQuery(updateData, map[string]string{
		"title":             "string",
		"description":       "string",
		"location":          "string",
		"enabled":           "bool",
		"enable_updated_at": "custom",
		"interval":          "int",
		"expected_status":   "int",
		"std_response_time": "int",
		"max_response_time": "int",
		"retain_duration":   "int",
		"hook_level":        "int",
		"hook_addr":         "string",
		"hook_secret":       "string",
	}, "uptime_watch_request", closing)

	if err != nil {
		return i, err
	}

	row := q.db.QueryRowContext(ctx, qry, time.Now().UTC(), id)
	err = row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Location,
		&i.Enabled,
		&i.EnableUpdatedAt,
		&i.Interval,
		&i.ExpectedStatus,
		&i.StdResponseTime,
		&i.MaxResponseTime,
		&i.RetainDuration,
		&i.HookLevel,
		&i.HookAddress,
	)
	return i, err
}
