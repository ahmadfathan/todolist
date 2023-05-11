package entity

import "errors"

var (
	ErrNotFound                   = errors.New("error no data found")
	ErrTitleCanotBeNull           = errors.New("title cannot be null")
	ErrActivityGroupIDCanotBeNull = errors.New("activity_group_id cannot be null")
)
