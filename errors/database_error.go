package errors

import "errors"

var ErrConnectionFailed = errors.New("failed to connect to database")
var ErrTablesCreationFailed = errors.New("failes to create tables")
var ErrExecuteFailed = errors.New("execute statement failed")
