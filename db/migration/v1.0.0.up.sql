CREATE TABLE IF NOT EXISTS key_value_store (
    key_data TEXT PRIMARY KEY,
    value_data TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
CREATE TABLE IF NOT EXISTS uptime_watch_request(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    enabled BOOLEAN NOT NULL,
    enable_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    interval INTEGER NOT NULL,
    expected_status INTEGER NOT NULL,
    max_response_time INTEGER NOT NULL,
    retain_duration INTEGER NOT NULL,
    hook_level INTEGER NOT NULL,
    hook_addr TEXT NOT NULL,
    hook_secret TEXT NOT NULL
);