package db

import (
	"context"
	"time"
)

const addKeyValue = `
INSERT INTO key_value_store(key_data, value_data, updated_at)
VALUES(?,?,?)
RETURNING key_data, value_data, updated_at
`

func (q *Queries) AddKeyValue(ctx context.Context, arg KeyValue) (KeyValue, error) {
	row := q.db.QueryRowContext(ctx, addKeyValue, arg.Key, arg.Value, time.Now())
	var i KeyValue
	err := row.Scan(
		&i.Key,
		&i.Value,
		&i.UpdateAt,
	)
	return i, err
}

const updateKeyValue = `
UPDATE key_value_store SET value_data = ? WHERE key_data = ?
RETURNING key_data, value_data, updated_at
`

func (q *Queries) UpdateKeyValue(ctx context.Context, arg KeyValue) (KeyValue, error) {
	row := q.db.QueryRowContext(ctx, updateKeyValue, arg.Value, arg.Key, time.Now())
	var i KeyValue
	err := row.Scan(
		&i.Key,
		&i.Value,
		&i.UpdateAt,
	)
	return i, err
}

const getKeyValue = `
SELECT key_data, value_data, updated_at FROM key_value_store
WHERE key_data = ?
`

func (q *Queries) GetKeyValue(ctx context.Context, key string) (KeyValue, error) {
	row := q.db.QueryRowContext(ctx, getKeyValue, key)
	var i KeyValue
	err := row.Scan(
		&i.Key,
		&i.Value,
		&i.UpdateAt,
	)
	return i, err
}

const deleteKeyValue = `
DELETE FROM key_value_store
WHERE key_data = ?
`

func (q *Queries) DeleteKeyValue(ctx context.Context, key string) error {
	_, err := q.db.ExecContext(ctx, deleteKeyValue, key)
	return err
}
