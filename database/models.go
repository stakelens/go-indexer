// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
)

type FetchLogsRangeCache struct {
	ID   string
	Data string
}

type RocketpoolTvl struct {
	ID          int64
	EthLocked   string
	RplLocked   string
	BlockNumber sql.NullInt64
}