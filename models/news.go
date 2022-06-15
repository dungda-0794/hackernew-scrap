package models

import "time"

type News struct {
	Id         string
	IdExternal string
	Title      string
	Author     string
	Link       string
	Point      int
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
