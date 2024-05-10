package entity

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/take0fit/knowledge-out/internal/domain/valueobject"
	"time"
)

type User struct {
	Id           string
	GoogleUserId string
	ImageUrl     string
	AccessToken  string
	RefreshToken string
	TokenType    string
	Expiry       string
	Nickname     valueobject.NickName
	LastName     string
	FirstName    string
	Birthday     valueobject.Birthday
	Gender       valueobject.Gender
	CreatedAt    string
	UpdatedAt    string
}

func NewUser(nickname string, birthday *string, gender *string) (*User, error) {
	userId := fmt.Sprintf("User#%s", uuid.New().String())
	nicknameObj, err := valueobject.NewUserNickname(nickname)
	if err != nil {
		return nil, err
	}
	genderObj, err := valueobject.NewGender(gender)
	if err != nil {
		return nil, err
	}

	newUser := User{
		Id:        userId,
		Nickname:  nicknameObj,
		Birthday:  valueobject.NewBirthday(birthday),
		Gender:    genderObj,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	return &newUser, nil
}

func MapGoogleUserToUser(googleUser *GoogleUser) (*User, error) {
	userId := fmt.Sprintf("User#%s", uuid.New().String())
	if googleUser.Name == "" {
		googleUser.Name = "noname"
	}
	nicknameObj, err := valueobject.NewUserNickname(googleUser.Name)
	if err != nil {
		return nil, err
	}
	genderObj, err := valueobject.NewGender(&googleUser.Gender)
	if err != nil {
		return nil, err
	}
	return &User{
		Id:           userId,
		GoogleUserId: googleUser.Sub,
		ImageUrl:     googleUser.Picture,
		//FirstName:    googleUser.FirstName,
		//LastName:     googleUser.LastName,
		Nickname:  nicknameObj,
		Gender:    genderObj,
		Birthday:  valueobject.NewBirthday(&googleUser.Birthdate),
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

func (b User) GetBirthday() *time.Time {
	if b.Birthday.Valid {
		return &b.Birthday.Time
	}

	return nil
}
