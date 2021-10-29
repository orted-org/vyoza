package db

import "time"

type KeyValue struct {
	Key      string    `json:"key"`
	Value    string    `json:"value"`
	UpdateAt time.Time `json:"updated_at"`
}

type UptimeWatchRequest struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Location        string    `json:"location"`
	Enabled         bool      `json:"enabled"`
	EnableUpdatedAt time.Time `json:"enable_updated_at"`

	// in seconds
	Interval   int  `json:"interval"`
	SSLMonitor bool `json:"ssl_monitor"`

	// in seconds
	SSLInterval int `json:"ssl_interval"`

	// in hours
	SSLExpiryNotification int `json:"ssl_expiry_notification"`
	ExpectedStatus        int `json:"expected_status"`

	// in milliseconds
	StdResponseTime int `json:"std_response_time"`

	// in milliseconds
	MaxResponseTime int `json:"max_response_time"`

	// in hours
	RetainDuration    int    `json:"retain_duration"`
	HookLevel         int    `json:"hook_level"`
	HookAddress       string `json:"hook_address"`
	NotificationEmail string `json:"notification_email"`
	HookSecret        string `json:"-"`
}

type UptimeResult struct {
	UWRID int `json:"uwr_id"`

	// in milliseconds
	ResponseTime int       `json:"response_time"`
	Remark       string    `json:"remark"`
	CreatedAt    time.Time `json:"created_at"`
}

type UptimeResultStats struct {
	UWRID        int `json:"uwr_id"`
	SuccessCount int `json:"success_count"`
	WarningCount int `json:"warning_count"`
	ErrorCount   int `json:"error_count"`

	// in milliseconds
	MinResponseTime int `json:"min_response_time"`

	// in milliseconds
	MaxResponseTime int `json:"max_response_time"`

	// in milliseconds
	AvgSuccessResponseTime int `json:"avg_success_resp_time"`

	// in milliseconds
	AvgWarningResponseTime int       `json:"avg_warning_resp_time"`
	StartDate              time.Time `json:"start_date"`
	EndDate                time.Time `json:"end_date"`
}

type UptimeConclusion struct {
	UWRID                  int       `json:"uwr_id"`
	SuccessCount           int       `json:"success_count"`
	WarningCount           int       `json:"warning_count"`
	ErrorCount             int       `json:"error_count"`
	MinResponseTime        int       `json:"min_response_time"`
	MaxResponseTime        int       `json:"max_response_time"`
	AvgSuccessResponseTime int       `json:"avg_success_resp_time"`
	AvgWarningResponseTime int       `json:"avg_warning_resp_time"`
	StartDate              time.Time `json:"start_date"`
	EndDate                time.Time `json:"end_date"`
}

type UptimeSSLInfo struct {
	UWRID      int       `json:"uwr_id"`
	IsValid    bool      `json:"is_valid"`
	ExpiryDate time.Time `json:"expiry_date"`
	Remark     string    `json:"remark"`
	UpdatedAt  time.Time `json:"updated_at"`
}
