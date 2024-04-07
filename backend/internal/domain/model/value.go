package model

import "errors"

type UserName string

func NewUserName(name string) (*UserName, error) {
	if len(name) < 1 || len(name) > 10 {
		return nil, errors.New("名前は1文字以上10文字以内で入力してください")
	}
	n := UserName(name)
	return &n, nil
}
