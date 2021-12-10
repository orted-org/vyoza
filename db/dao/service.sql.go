package db

import (
	"context"
	"time"
)

const addService = `
INSERT INTO service (
	id,
	name,
	description,
	secret,
	created_at
)
VALUES (?, ?, ?, ?, ?)
RETURNING id,
	name,
	description,
	secret,
	created_at;
`

func (q *Queries) AddService(ctx context.Context, arg Service) (Service, error) {
	arg.CreatedAt = time.Now().UTC()
	row := q.queryRow(
		ctx,
		q.addService,
		addService,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Secret,
		arg.CreatedAt,
	)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Secret,
		&i.CreatedAt,
	)
	return i, err
}

const deleteServiceByID = `
DELETE FROM service
WHERE id = ?;
`

func (q *Queries) DeleteServiceByID(ctx context.Context, ID string) error {
	_, err := q.exec(ctx, q.deleteServiceByID, deleteServiceByID, ID)
	return err
}

const getServiceByID = `
SELECT id,
    name,
    description,
    secret,
    created_at
FROM service
WHERE id = ?;
`

func (q *Queries) GetServiceByID(ctx context.Context, ID string) (Service, error) {
	row := q.queryRow(
		ctx,
		q.getServiceByID,
		getServiceByID,
		ID,
	)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Secret,
		&i.CreatedAt,
	)
	return i, err
}

const getAllService = `
SELECT id,
    name,
    description,
    secret,
    created_at
FROM service;
`

func (q *Queries) GetAllService(ctx context.Context) ([]Service, error) {
	rows, err := q.query(ctx, q.getAllService, getAllService)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Service
	for rows.Next() {
		var i Service
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Secret,
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
