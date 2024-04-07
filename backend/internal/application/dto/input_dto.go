package dto

type InputCreateInput struct {
	UserId     string
	ResourceId string
	Name       string
	Detail     string
	CategoryId int
}

type InputOutput struct {
	Id         string
	UserId     string
	ResourceId string
	Name       string
	Detail     string
	CategoryId int
}
