package model

import "time"

// Output は、学んだことをどのように活用したかを表します。
type Output struct {
	ID        string    `json:"id"`        // アウトプットの一意識別子
	InputID   string    `json:"inputId"`   // 関連するインプットのID
	Action    string    `json:"action"`    // 実際に取った行動
	Result    string    `json:"result"`    // 行動の結果
	CreatedAt time.Time `json:"createdAt"` // 行動を起こした日時
}
