package valueobject

import (
	"errors"
	"log"
)

type NickName string

func NewUserNickname(nickname string) (NickName, error) {
	if len(nickname) < 1 || len(nickname) > 10 {
		log.Printf("nicname: %s", nickname)

		return "", errors.New("ニックネームは1文字以上10文字以内で入力してください")
	}
	return NickName(nickname), nil
}

func (n NickName) String() string {

	return string(n)
}
