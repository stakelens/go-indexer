CREATE TABLE IF NOT EXISTS rocketpool_tvl (
    id SERIAL PRIMARY KEY,
    eth_locked TEXT NOT NULL,
    rpl_locked TEXT NOT NULL,
    block_number BIGINT
);

CREATE TABLE IF NOT EXISTS cache (
    id TEXT PRIMARY KEY,
    data TEXT NOT NULL
);
