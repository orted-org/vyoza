-- create key value pair
INSERT INTO key_value_store(key_data, value_data, updated_at)
VALUES(?, ?, ?)
RETURNING key_data,
    value_data,
    updated_at;
-- update key value
UPDATE key_value_store
SET value_data = ?
WHERE key_data = ?
RETURNING key_data,
    value_data,
    updated_at;
-- get key value
SELECT key_data,
    value_data,
    updated_at
FROM key_value_store
WHERE key_data = ?;
-- delete key value
DELETE FROM key_value_store
WHERE key_data = ?