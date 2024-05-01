-- name: GetAllRocketPoolTVLs :many
SELECT * FROM rocketpool_tvl ORDER BY block_number DESC;

-- name: SaveRocketPoolTVL :exec
INSERT INTO rocketpool_tvl (eth_locked, rpl_locked, block_number) VALUES (?, ?, ?);

-- name: Cache :exec
INSERT INTO cache (id, data) VALUES (?, ?);

-- name: GetCache :one
SELECT * FROM cache WHERE id = ?;
