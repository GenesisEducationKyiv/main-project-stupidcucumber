package defined_errors

import "errors"

var (
	EmptyCache = errors.New("empty cache")
)
