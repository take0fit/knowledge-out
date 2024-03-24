package model

type User struct {
	ID   string `dynamodbav:"Id"`
	Name string `dynamodbav:"Name"`
	Age  int    `dynamodbav:"Age"`
}
