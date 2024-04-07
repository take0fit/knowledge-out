package dto

type ResourceCreateInput struct {
	UserId     string
	Name       string
	Detail     string
	CategoryId int
}

type ResourceOutput struct {
	Id         string
	UserId     string
	Name       string
	Detail     string
	CategoryId int
}
