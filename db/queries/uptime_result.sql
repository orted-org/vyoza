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