--name : AddUptimeConclusion
INSERT INTO uptime_conclusion (
    id,
    success_count,
    warning_count,
    error_count,
    min_response_time,
    max_response_time,
    avg_success_resp_time,
    avg_warning_resp_time,
    success_count,
    start_date,
    end_date
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id,
    success_count,
    warning_count,
    error_count,
    min_response_time,
    max_response_time,
    avg_success_resp_time,
    avg_warning_resp_time,
    success_count,
    start_date,
    end_date;

--name : DeleteUptimeConclusionByID
DELETE FROM uptime_conclusion
WHERE id = ?;

--name : GetUptimeConclusionByID
SELECT id,
    success_count,
    warning_count,
    error_count,
    min_response_time,
    max_response_time,
    avg_success_resp_time,
    avg_warning_resp_time,
    success_count,
    start_date,
    end_date
FROM uptime_conclusion;