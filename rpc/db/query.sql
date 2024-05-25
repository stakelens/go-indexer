-- name: Cache :exec
INSERT INTO cache (id, data) VALUES ($1, $2);

-- name: GetCache :one
SELECT * FROM cache WHERE id = $1;
