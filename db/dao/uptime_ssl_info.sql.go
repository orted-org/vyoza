package db

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

const deleteUptimeSSLInfoByUWRID = `
DELETE FROM 
    uptime_ssl_info
WHERE 
    uwr_id = ?;
`

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

const getAllUptimeSSLInfo = `
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
`


