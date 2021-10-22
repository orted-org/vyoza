--name : AddUptimeSSLInfo
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

--name : DeleteUptimeSSLInfoByUWRID
DELETE FROM 
    uptime_ssl_info
WHERE 
    uwr_id = ?;


--name : UpdateUptimeSSLInfoByUWRID
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

--name : GetUptimeSSLInfoByUWRID
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


--name : GetAllUptimeSSLInfo
SELECT
    uwr_id,
    is_valid,
    expiry_date,
    remark,
    updated_at
FROM 
    uptime_ssl_info;
LIMIT ?
OFFSET ?;

