package db

import (
	"context"
)

const addUptimeConclusion = `
INSERT INTO uptime_conclusion (
    uwr_id,
    success_count,
    warning_count,
    error_count,
    min_response_time,
    max_response_time,
    avg_success_resp_time,
    avg_warning_resp_time,
    start_date,
    end_date
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING uwr_id,
    success_count,
    warning_count,
    error_count,
    min_response_time,
    max_response_time,
    avg_success_resp_time,
    avg_warning_resp_time,
    start_date,
    end_date;
`

func (q *Queries) AddUptimeConclusion(ctx context.Context, arg UptimeConclusion) (UptimeConclusion, error) {
	row := q.queryRow(
		ctx,
		q.addUptimeConclusion,
		addUptimeConclusion,
		arg.UWRID,
		arg.SuccessCount,
		arg.WarningCount,
		arg.ErrorCount,
		arg.MinResponseTime,
		arg.MaxResponseTime,
		arg.AvgSuccessResponseTime,
		arg.AvgWarningResponseTime,
		arg.StartDate,
		arg.EndDate,
	)
	var i UptimeConclusion
	err := row.Scan(
		&i.UWRID,
		&i.SuccessCount,
		&i.WarningCount,
		&i.ErrorCount,
		&i.MinResponseTime,
		&i.MaxResponseTime,
		&i.AvgSuccessResponseTime,
		&i.AvgWarningResponseTime,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const deleteUptimeConclusionByUWRID = `
DELETE FROM uptime_conclusion
WHERE uwr_id = ?;
`

func (q *Queries) DeleteUptimeConclusionByUWRID(ctx context.Context, uwr_id int) error {
	_, err := q.exec(ctx, q.deleteUptimeConclusionByUWRID, deleteUptimeConclusionByUWRID, uwr_id)
	return err
}

const getUptimeConclusionByUWRID = `
SELECT uwr_id,
    success_count,
    warning_count,
    error_count,
    min_response_time,
    max_response_time,
    avg_success_resp_time,
    avg_warning_resp_time,
    start_date,
    end_date
FROM uptime_conclusion
WHERE uwr_id = ?
`

func (q *Queries) GetUptimeConclusionByUWRID(ctx context.Context, uwr_id int) (UptimeConclusion, error) {
	row := q.queryRow(ctx, q.getUptimeConclusionByUWRID, getUptimeConclusionByUWRID, uwr_id)
	var i UptimeConclusion
	err := row.Scan(
		&i.UWRID,
		&i.SuccessCount,
		&i.WarningCount,
		&i.ErrorCount,
		&i.MinResponseTime,
		&i.MaxResponseTime,
		&i.AvgSuccessResponseTime,
		&i.AvgWarningResponseTime,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const getAllUptimeConclusion = `
SELECT uwr_id,
    success_count,
    warning_count,
    error_count,
    min_response_time,
    max_response_time,
    avg_success_resp_time,
    avg_warning_resp_time,
    start_date,
    end_date
FROM uptime_conclusion
LIMIT ?
OFFSET ?
;
`

type getAllUptimeConclusionParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func (q *Queries) GetAllUptimeConclusion(ctx context.Context, arg getAllUptimeConclusionParams) ([]UptimeConclusion, error) {
	rows, err := q.query(ctx, q.getAllUptimeConclusion, getAllUptimeConclusion, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UptimeConclusion
	for rows.Next() {
		var i UptimeConclusion
		if err := rows.Scan(
			&i.UWRID,
			&i.SuccessCount,
			&i.WarningCount,
			&i.ErrorCount,
			&i.MinResponseTime,
			&i.MaxResponseTime,
			&i.AvgSuccessResponseTime,
			&i.AvgWarningResponseTime,
			&i.StartDate,
			&i.EndDate,
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
