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
    ssl_interval,
    ssl_expiry_notification,
    expected_status,
    std_response_time,
    max_response_time,
    retain_duration,
    hook_level,
    hook_addr,
    hook_secret,
    notification_email
)
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	RETURNING id,
	title,
	description,
	location,
	enabled,
	enable_updated_at,
	interval,
	ssl_monitor,
	ssl_interval,
    ssl_expiry_notification,
	expected_status,
	std_response_time,
	max_response_time,
	retain_duration,
	hook_level,
	hook_addr,
	hook_secret,
	notification_email
`

type AddUptimeWatchRequestParams struct {
	Title                 string `json:"title"`
	Description           string `json:"description"`
	Location              string `json:"location"`
	Enabled               bool   `json:"enabled"`
	Interval              int    `json:"interval"`
	SSLMonitor            bool   `json:"ssl_monitor"`
	SSLInterval           int    `json:"ssl_interval"`
	SSLExpiryNotification int    `json:"ssl_expiry_notification"`
	ExpectedStatus        int    `json:"expected_status"`
	StdResponseTime       int    `json:"std_response_time"`
	MaxResponseTime       int    `json:"max_response_time"`
	RetainDuration        int    `json:"retain_duration"`
	HookLevel             int    `json:"hook_level"`
	HookAddress           string `json:"hook_address"`
	HookSecret            string `json:"hook_secret"`
	NotificationEmail     string `json:"notification_email"`
}

func (q *Queries) AddUptimeWatchRequest(ctx context.Context, arg AddUptimeWatchRequestParams) (UptimeWatchRequest, error) {
	row := q.queryRow(ctx, q.addUptimeWatchRequest, addUptimeWatchRequest, arg.Title, arg.Description, arg.Location, arg.Enabled, time.Now().UTC(), arg.Interval, arg.SSLMonitor, arg.SSLInterval, arg.SSLExpiryNotification, arg.ExpectedStatus, arg.StdResponseTime, arg.MaxResponseTime, arg.RetainDuration, arg.HookLevel, arg.HookAddress, arg.HookSecret, arg.NotificationEmail)
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
		&i.SSLInterval,
		&i.SSLExpiryNotification,
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
    ssl_monitor,
	ssl_interval,
    ssl_expiry_notification,
    expected_status,
    std_response_time,
    max_response_time,
    retain_duration,
    hook_level,
    hook_addr,
    hook_secret,
    notification_email
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
		&i.SSLMonitor,
		&i.SSLInterval,
		&i.SSLExpiryNotification,
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

const getAllUptimeWatchRequest = `
SELECT id,
    title,
    description,
    location,
    enabled,
    enable_updated_at,
    interval,
    ssl_monitor,
	ssl_interval,
    ssl_expiry_notification,
    expected_status,
    std_response_time,
    max_response_time,
    retain_duration,
    hook_level,
    hook_addr,
    hook_secret,
    notification_email
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
			&i.SSLMonitor,
			&i.SSLInterval,
			&i.SSLExpiryNotification,
			&i.ExpectedStatus,
			&i.StdResponseTime,
			&i.MaxResponseTime,
			&i.RetainDuration,
			&i.HookLevel,
			&i.HookAddress,
			&i.HookSecret,
			&i.NotificationEmail,
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
		ssl_monitor,
		ssl_interval,
    	ssl_expiry_notification,
		expected_status,
		std_response_time,
		max_response_time,
		retain_duration,
		hook_level,
		hook_addr,
		notification_email
	`
	qry, err := CreateDynamicUpdateQuery(updateData, map[string]string{
		"title":                   "string",
		"description":             "string",
		"location":                "string",
		"enabled":                 "bool",
		"enable_updated_at":       "custom",
		"interval":                "int",
		"ssl_monitor":             "bool",
		"ssl_interval":            "int",
		"ssl_expiry_notification": "int",
		"expected_status":         "int",
		"std_response_time":       "int",
		"max_response_time":       "int",
		"retain_duration":         "int",
		"hook_level":              "int",
		"hook_addr":               "string",
		"hook_secret":             "string",
		"notification_email":      "string",
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
		&i.SSLMonitor,
		&i.SSLInterval,
		&i.SSLExpiryNotification,
		&i.ExpectedStatus,
		&i.StdResponseTime,
		&i.MaxResponseTime,
		&i.RetainDuration,
		&i.HookLevel,
		&i.HookAddress,
		&i.NotificationEmail,
	)
	return i, err
}
