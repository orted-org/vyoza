CREATE TABLE IF NOT EXISTS key_value_store (
    key_data TEXT PRIMARY KEY,
    value_data TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS uptime_watch_request(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    enabled BOOLEAN NOT NULL,
    interval INTEGER NOT NULL,
    expected_status INTEGER NOT NULL,
    retain_duration TEXT,
    hook_level INTEGER,
    hook_addr TEXT
);