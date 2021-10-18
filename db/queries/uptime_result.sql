-- create new result 
INSERT INTO uptime_result(id, response_time)
VALUES(?, ?)
RETURNING id,
    response_time,
    created_at;
-- get count of the result by id
SELECT count(*)
WHERE id = ?;
-- get all the uptime values for the specified limit and range
SELECT id,
    response_time,
    created_at
FROM uptime_result
WHERE id = ?
LIMIT ? OFFSET ?;
-- delete uptime result 
DELETE FROM uptime_result
WHERE id = ?;
-- get stats for an id
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
            WHEN response_time > uptime_watch_request.max_response_time THEN response_time
        END
    ) AS error_count,
    min(response_time) AS min_response_time,
    max(response_time) AS max_response_time,
    avg(
        CASE
            WHEN response_time <= uptime_watch_request.std_response_time THEN response_time
        END
    ) AS avg_success_resp_time,
    avg(
        CASE
            WHEN response_time > uptime_watch_request.std_response_time
            AND response_time <= uptime_watch_request.max_response_time THEN response_time
        END
    ) AS avg_warning_resp_time,
    min(created_at) AS start_date,
    max(created_at) AS end_date
FROM uptime_result
    INNER JOIN uptime_watch_request ON uptime_watch_request.id = uptime_result.id
WHERE uptime_watch_request.id = ?;