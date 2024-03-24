package dto

type UserCreateInput struct {
	Name string
	Age  int
}

type UserOutput struct {
	ID   string
	Name string
	Age  int
}
