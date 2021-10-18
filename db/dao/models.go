package db

import "time"

type KeyValue struct {
	Key      string    `json:"key"`
	Value    string    `json:"value"`
	UpdateAt time.Time `json:"updated_at"`
}

/*
	Location -> URL of the service
	Enabled -> whether watch is enabled
	EnableUpdatedAt -> timestamp of change in value of enabled
	Interval -> watch time interval(in seconds)
	ExpectedStatus -> expected status code for the service to return
	MaxResponeTime -> if status code matched and within this time, then considered success, otherwise warning
	RetainDuration -> duration to retain the data in db (1h, 3d, etc)
	HookLevel -> to make hook HTTP request at event 1(for only error), 2(for error and warning), 3(for success, warning and error)
	HookAddress -> URL to make post hook request
	HookSecret -> secret to send to the hook call to make safe hook call
*/
type UptimeWatchRequest struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Location        string    `json:"location"`
	Enabled         bool      `json:"enabled"`
	EnableUpdatedAt time.Time `json:"enable_updated_at"`
	Interval        int       `json:"interval"`
	ExpectedStatus  int       `json:"expected_status"`
	StdResponseTime int       `json:"std_response_time"`
	MaxResponseTime int       `json:"max_response_time"`
	RetainDuration  int       `json:"retain_duration"`
	HookLevel       int       `json:"hook_level"`
	HookAddress     string    `json:"hook_address"`
	HookSecret      string
}

type UptimeResult struct {
	ID           int       `json:"id"`
	ResponseTime int       `json:"response_time"`
	CreatedAt    time.Time `json:"created_at"`
}
