package constant

import "errors"

var (
	ErrNoRecordPresent = errors.New("sql: no rows in result set")
)
