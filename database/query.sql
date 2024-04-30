-- name: GetAllRocketPoolTVLs :many
SELECT * FROM rocketpool_tvl ORDER BY block_number DESC;

-- name: SaveRocketPoolTVL :exec
INSERT INTO rocketpool_tvl (eth_locked, rpl_locked, block_number) VALUES (?, ?, ?);

-- name: CacheFetchLogsRange :exec
INSERT INTO fetch_logs_range_cache (id, data) VALUES (?, ?);

-- name: GetCachedFetchLogsRange :one
SELECT * FROM fetch_logs_range_cache WHERE id = ?;
