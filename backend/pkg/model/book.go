package model

import "time"

// Book は、読んだ本の情報を表します。
type Book struct {
	ID          string    `json:"id"`          // 本の一意識別子
	Title       string    `json:"title"`       // 本のタイトル
	Author      string    `json:"author"`      // 著者名
	Url         string    `json:"url"`         // url
	PublishedAt time.Time `json:"publishedAt"` // 出版日
}
