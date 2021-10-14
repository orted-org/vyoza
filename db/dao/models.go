package db

import "time"

type KeyValue struct {
	Key      string    `json:"key"`
	Value    string    `json:"value"`
	UpdateAt time.Time `json:"updated_at"`
}
type UptimeWatchRequest struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Location       string `json:"location"`
	Enabled        bool   `json:"enabled"`
	Interval       int    `json:"interval"`
	ExpectedStatus int    `json:"expected_status"`
	RetainDuration string `json:"retain_duration"`
	HookLevel      int    `json:"hook_level"`
	HookAddress    string `json:"hook_address"`
}
