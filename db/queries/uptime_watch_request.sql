-- insert uptime watch request
INSERT INTO uptime_watch_request (
        INSERT INTO uptime_watch_request (
                title,
                description,
                location,
                enabled,
                enable_updated_at,
                interval,
                expected_status,
                max_response_time,
                retain_duration,
                hook_level,
                hook_addr,
                hook_secret
            )
        VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        RETURNING id,
            title,
            description,
            location,
            enabled,
            enable_updated_at,
            interval,
            expected_status,
            max_response_time,
            retain_duration,
            hook_level,
            hook_addr,
            hook_secret;
-- get one uptime watch request by id
SELECT id,
    title,
    description,
    location,
    enabled,
    enable_updated_at,
    interval,
    expected_status,
    max_response_time,
    retain_duration,
    hook_level,
    hook_addr,
    hook_secret
FROM uptime_watch_request
WHERE id = ?;
-- get all uptime watch request
SELECT id,
    title,
    description,
    location,
    enabled,
    enable_updated_at,
    interval,
    expected_status,
    max_response_time,
    retain_duration,
    hook_level,
    hook_addr,
    hook_secret
FROM uptime_watch_request;