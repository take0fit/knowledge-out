package model

import "time"

// Input は、本から学んだことを表します。
type Input struct {
	ID        string    `json:"id"`        // インプットの一意識別子
	BookID    string    `json:"bookId"`    // 関連する本のID
	Content   string    `json:"content"`   // 学んだ内容
	CreatedAt time.Time `json:"createdAt"` // 学んだ日時
}
