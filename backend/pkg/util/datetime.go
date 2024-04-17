package util

import "time"

// IsBefore compares two time.Time objects and returns true if the first one
// is before the second one, ignoring their time of the day.
// IsBefore は2つの time.Time オブジェクトを比較し、第一引数の日時が
// 第二引数の日時よりも前であるかどうかを返します。時間は無視され、日付のみが考慮されます。
func IsBefore(first, second time.Time) bool {
	if first.Year() < second.Year() {
		return true
	}
	if first.Year() > second.Year() {
		return false
	}
	// 年が同じ場合、月と日で判断
	if first.Month() < second.Month() {
		return true
	}
	if first.Month() > second.Month() {
		return false
	}
	return first.Day() < second.Day()
}
