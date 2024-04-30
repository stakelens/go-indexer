CREATE TABLE rocketpool_tvl (
    id INTEGER PRIMARY KEY,
    eth_locked TEXT NOT NULL,
    rpl_locked TEXT NOT NULL,
    block_number BIGINT
);
