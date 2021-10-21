CREATE TABLE IF NOT EXISTS key_value_store (
    key_data TEXT PRIMARY KEY,
    value_data TEXT NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
CREATE TABLE IF NOT EXISTS uptime_watch_request(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    enabled BOOLEAN NOT NULL,
    enable_updated_at TIMESTAMP NOT NULL,
    interval INTEGER NOT NULL,
    ssl_monitor BOOLEAN NOT NULL,
    expected_status INTEGER NOT NULL,
    std_response_time INTEGER NOT NULL,
    max_response_time INTEGER NOT NULL,
    retain_duration INTEGER NOT NULL,
    hook_level INTEGER NOT NULL,
    hook_addr TEXT NOT NULL,
    hook_secret TEXT NOT NULL,
    notification_email TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS uptime_result(
    id INTEGER REFERENCES uptime_watch_request(id),
    response_time INTEGER NOT NULL,
    remark TEXT NOT NULL,
    created_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS uptime_conclusion (
    uwr_id INTEGER UNIQUE REFERENCES uptime_watch_request(id),
    success_count INTEGER NOT NULL,
    warning_count INTEGER NOT NULL,
    error_count INTEGER NOT NULL,
    min_response_time INTEGER NOT NULL,
    max_response_time INTEGER NOT NULL,
    avg_success_resp_time INTEGER NOT NULL,
    avg_warning_resp_time INTEGER NOT NULL,
    start_date TIMESTAMP,
    end_date TIMESTAMP
)