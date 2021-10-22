package db

import (
	"context"
)

const addUptimeSSLInfo = `
INSERT INTO 
uptime_ssl_info (
    uwr_id,
    is_valid,
    expiry_date,
    remark,
    updated_at
) 
VALUES 
    (?, ?, ?, ?, ?)
RETURNING 
    uwr_id,
    is_valid,
    expiry_date,
    remark,
    updated_at;
`

func (q *Queries) AddUptimeSSLInfo(ctx context.Context, arg UptimeSSLInfo) (UptimeSSLInfo, error) {
	row := q.queryRow(
        ctx,
        q.addUptimeSSLInfo,
        addUptimeSSLInfo,
        arg.UWRId,
        arg.IsValid,
        arg.ExpiryDate,
        arg.Remark,
        arg.UpdatedAt,
	)
	var i UptimeSSLInfo
	err := row.Scan(
		&i.UWRId,
		&i.IsValid,
		&i.ExpiryDate,
		&i.Remark,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUptimeSSLInfoByUWRID = `
DELETE FROM 
    uptime_ssl_info
WHERE 
    uwr_id = ?;
`
func (q *Queries) DeleteUptimeSSLInfoByUWRID(ctx context.Context, uwr_id int) error {
	_, err := q.exec(ctx, q.deleteUptimeSSLInfoByUWRID, deleteUptimeSSLInfoByUWRID, uwr_id)
	return err
}

const updateUptimeSSLInfoByUWRID = `
UPDATE 
    uptime_ssl_info
SET
    is_valid = ?,
    expiry_date = ?,
    remark = ?,
    updated_at = ?
WHERE 
    uwr_id = ?
RETURNING
    uwr_id,
    is_valid,
    expiry_date,
    remark,
    updated_at;
`


func (q *Queries) UpdateUptimeSSLInfoByUWRID(ctx context.Context, arg UptimeSSLInfo) (UptimeSSLInfo, error) {
	row := q.queryRow(
        ctx,
        q.updateUptimeSSLInfoByUWRID,
        updateUptimeSSLInfoByUWRID,
        arg.IsValid,
        arg.ExpiryDate,
        arg.Remark,
        arg.UpdatedAt,
        arg.UWRId,
	)
	var i UptimeSSLInfo
	err := row.Scan(
		&i.UWRId,
		&i.IsValid,
		&i.ExpiryDate,
		&i.Remark,
		&i.UpdatedAt,
	)
	return i, err
}

const getUptimeSSLInfoByUWRID = `
SELECT
    uwr_id,
    is_valid,
    expiry_date,
    remark,
    updated_at
FROM 
    uptime_ssl_info
WHERE
    uwr_id = ?;
`
func (q *Queries) GetUptimeSSLInfoByUWRID(ctx context.Context, uwr_id int) (UptimeSSLInfo, error) {
	row := q.queryRow(
        ctx,
        q.getUptimeSSLInfoByUWRID,
        getUptimeSSLInfoByUWRID,
        uwr_id,
	)
	var i UptimeSSLInfo
	err := row.Scan(
		&i.UWRId,
		&i.IsValid,
		&i.ExpiryDate,
		&i.Remark,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllUptimeSSLInfo = `
SELECT
    uwr_id,
    is_valid,
    expiry_date,
    remark,
    updated_at
FROM 
    uptime_ssl_info
LIMIT ?
OFFSET ?;
`
type getAllUptimeSSLInfoParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func (q *Queries) GetAllUptimeSSLInfo(ctx context.Context, arg getAllUptimeSSLInfoParams) ([]UptimeSSLInfo, error) {
	rows, err := q.query(ctx, q.getAllUptimeSSLInfo, getAllUptimeSSLInfo, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UptimeSSLInfo
	for rows.Next() {
		var i UptimeSSLInfo
		if err := rows.Scan(
            &i.UWRId,
            &i.IsValid,
            &i.ExpiryDate,
            &i.Remark,
            &i.UpdatedAt,
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


