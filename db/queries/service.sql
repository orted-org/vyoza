-- name : AddService
INSERT INTO service (
        id,
        name,
        description,
        secret,
        created_at
    )
VALUES (?, ?, ?, ?, ?)
RETURNING id,
    name,
    description,
    secret,
    created_at;

-- name : DeleteServiceByID
DELETE FROM service
WHERE id = ?;

-- name : GetServiceByID
SELECT id,
    name,
    description,
    secret,
    created_at
FROM service
WHERE id = ?;

-- name : GetAllService
SELECT id,
    name,
    description,
    secret,
    created_at
FROM service;