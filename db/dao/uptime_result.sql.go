package db

import (
	"context"
)

const addUptimeResult = `
INSERT INTO uptime_result(id, response_time)
VALUES(?, ?)
RETURNING id,
    response_time,
    created_at
`

type AddUptimeResultParams struct {
	ID           int
	ResponseTime int
}

func (q *Queries) AddUptimeResult(ctx context.Context, arg AddUptimeResultParams) (UptimeResult, error) {
	row := q.queryRow(ctx, q.addUptimeResult, addUptimeResult, arg.ID, arg.ResponseTime)
	var i UptimeResult
	err := row.Scan(
		&i.ID,
		&i.ResponseTime,
		&i.CreatedAt,
	)
	return i, err
}

const getUptimeResultCount = `
SELECT count(*)
WHERE id = ?
`

func (q *Queries) GetUptimeResultCount(ctx context.Context, arg int) (int, error) {
	row := q.queryRow(ctx, q.getUptimeResultCount, getUptimeResultCount, arg)
	var i int
	err := row.Scan(
		&i,
	)
	return i, err
}

const getUptimeResults = `
SELECT id,
    response_time,
    created_at
FROM uptime_result
WHERE id = ?
LIMIT ? OFFSET ?
`

type GetUptimeResultsParams struct {
	ID     int
	Limit  int
	Offset int
}

func (q *Queries) GetUptimeResults(ctx context.Context, arg GetUptimeResultsParams) ([]UptimeResult, error) {
	rows, err := q.query(ctx, q.getUptimeResults, getUptimeResults, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UptimeResult
	for rows.Next() {
		var i UptimeResult
		if err := rows.Scan(
			&i.ID,
			&i.ResponseTime,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteUptimeResults = `
DELETE FROM uptime_result
WHERE id = ?
`

func (q *Queries) DeleteUptimeResults(ctx context.Context, id int) error {
	_, err := q.exec(ctx, q.deleteUptimeResults, deleteUptimeResults, id)
	return err
}
