package components

import (
	"kaonavi_driver/config"
	"time"
)

// 現在の時刻を取得する
func GetCurrentTime() string {
	// 現在の時刻を取得
	currentTime := time.Now()
	// フォーマットを指定して文字列に変換
	formattedTime := currentTime.Format(config.YYYYMMDDHHMMSS)

	return formattedTime
}
