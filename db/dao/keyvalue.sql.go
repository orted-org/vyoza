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
	row := q.queryRow(ctx, q.addKeyValue, addKeyValue, arg.Key, arg.Value, time.Now().UTC())
	var i KeyValue
	err := row.Scan(
		&i.Key,
		&i.Value,
		&i.UpdateAt,
	)
	return i, err
}

const updateKeyValue = `
UPDATE key_value_store SET value_data = ?, updated_at = ? WHERE key_data = ?
RETURNING key_data, value_data, updated_at
`

func (q *Queries) UpdateKeyValue(ctx context.Context, arg KeyValue) (KeyValue, error) {
	row := q.queryRow(ctx, q.updateKeyValue, updateKeyValue, arg.Value, time.Now().UTC(), arg.Key)
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
	row := q.queryRow(ctx, q.getKeyValue, getKeyValue, key)
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
	_, err := q.exec(ctx, q.deleteKeyValue, deleteKeyValue, key)
	return err
}
