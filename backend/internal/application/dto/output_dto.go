package dto

import "time"

type OutputCreateInput struct {
	UserId     string
	InputIds   []string
	Name       string
	Detail     string
	CategoryId int
}

type OutputOutput struct {
	Id         string
	UserId     string
	InputIds   []string
	Name       string
	Detail     string
	CategoryId int
	CreatedAt  time.Time
}
