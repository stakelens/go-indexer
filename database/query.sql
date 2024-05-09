-- name: GetAllRocketPoolTVLs :many
SELECT * FROM rocketpool_tvl ORDER BY block_number DESC;

-- name: SaveRocketPoolTVL :exec
INSERT INTO rocketpool_tvl (eth_locked, rpl_locked, block_number) VALUES ($1, $2, $3);

-- name: Cache :exec
INSERT INTO cache (id, data) VALUES ($1, $2);

-- name: GetCache :one
SELECT * FROM cache WHERE id = $1;

-- name: GetLatestBlockRocketPoolTVL :one
SELECT * FROM rocketpool_tvl ORDER BY block_number DESC LIMIT 1;