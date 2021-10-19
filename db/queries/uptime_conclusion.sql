--name : AddUptimeConclusion
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

--name : DeleteUptimeConclusionByUWRID
DELETE FROM uptime_conclusion
WHERE uwr_id = ?;

--name : GetUptimeConclusionByUWRID
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
FROM uptime_conclusion;
Where uwr_id = ?

--name : GetAllUptimeConclusion
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
FROM uptime_conclusion;