CREATE TABLE IF NOT EXISTS key_value_store (
    key_data TEXT PRIMARY KEY,
    value_data TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);