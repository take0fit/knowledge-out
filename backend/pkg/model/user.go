package model

// User は、アプリケーションのユーザーを表します。
type User struct {
	ID       string `json:"id"`       // ユーザーの一意識別子
	Username string `json:"username"` // ユーザー名
	Email    string `json:"email"`    // ユーザーのメールアドレス
}
