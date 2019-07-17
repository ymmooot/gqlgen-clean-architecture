package data

import "time"

type Article struct {
	ID          uint64
	Title       string
	Body        string
	PublishedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
