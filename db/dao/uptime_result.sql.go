package db

import (
	"context"
	"log"
	"strings"
	"time"
)

const addUptimeResult = `
INSERT INTO uptime_result(id, response_time, created_at)
VALUES(?, ?, ?)
RETURNING id,
    response_time,
    created_at
`

type AddUptimeResultParams struct {
	ID           int
	ResponseTime int
}

func (q *Queries) AddUptimeResult(ctx context.Context, arg AddUptimeResultParams) (UptimeResult, error) {
	row := q.queryRow(ctx, q.addUptimeResult, addUptimeResult, arg.ID, arg.ResponseTime, time.Now().UTC())
	var i UptimeResult
	err := row.Scan(
		&i.ID,
		&i.ResponseTime,
		&i.CreatedAt,
	)
	return i, err
}

const getUptimeResultCount = `
SELECT count(*) FROM uptime_result
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

const getUptimeResultStatsForID = `
SELECT uptime_result.id,
    count(
        CASE
            WHEN response_time <= uptime_watch_request.std_response_time THEN response_time
        END
    ) AS success_count,
    count(
        CASE
            WHEN response_time > uptime_watch_request.std_response_time
            AND response_time <= uptime_watch_request.max_response_time THEN response_time
        END
    ) AS warning_count,
    count(
        CASE
			WHEN response_time == -1 THEN response_time
        END
    ) AS error_count,
    min(response_time) AS min_response_time,
    max(response_time) AS max_response_time,
    CAST(
        IFNULL(
            avg(
                CASE
                    WHEN response_time <= uptime_watch_request.std_response_time THEN response_time
                END
            ),
            0
        ) AS INTEGER
    ) AS avg_success_resp_time,
    CAST(
        IFNULL(
            avg(
                CASE
                    WHEN response_time > uptime_watch_request.std_response_time
                    AND response_time <= uptime_watch_request.max_response_time THEN response_time
                END
            ),
            0
        ) AS INTEGER
    ) AS avg_warning_resp_time,
    min(created_at) AS start_date,
    max(created_at) AS end_date
FROM uptime_result
    INNER JOIN uptime_watch_request ON uptime_watch_request.id = uptime_result.id
WHERE uptime_watch_request.id = ?
`

func (q *Queries) GetUptimeResultStatsForID(ctx context.Context, id int) (UptimeResultStats, error) {
	row := q.queryRow(ctx, q.getUptimeResultStatsForID, getUptimeResultStatsForID, id)
	var i UptimeResultStats
	var startDate string
	var endDate string
	err := row.Scan(
		&i.ID,
		&i.SuccessCount,
		&i.WarningCount,
		&i.ErrorCount,
		&i.MinResponseTime,
		&i.MaxResponseTime,
		&i.AvgSuccessResponseTime,
		&i.AvgWarningResponseTime,
		&startDate,
		&endDate,
	)
	startDate = strings.Split(strings.ReplaceAll(startDate, " ", "T"), ".")[0] + ".000Z"
	endDate = strings.Split(strings.ReplaceAll(endDate, " ", "T"), ".")[0] + ".000Z"

	//incoming -> 2021-10-19 10:29:51.712726+05:45
	lay := "2006-01-02T15:04:05.000Z"

	i.StartDate, _ = time.Parse(lay, startDate)
	if err != nil {
		log.Println(err)
	}
	i.EndDate, _ = time.Parse(lay, endDate)

	return i, err
}
