package config

const (
	// time
	// YYYY-MM-DD HH:MM:SS
	YYYYMMDDHHMMSS = "2006-01-02 15:04:05"

	// EventType
	// メンバーが新しく登録された
	EVENT_TYPE_MEMBER_CREATED = "member_created"
	// メンバーの基本情報が更新された
	EVENT_TYPE_MEMBER_UPDATED = "member_updated"
	// メンバーが削除された
	EVENT_TYPE_MEMBER_DELETED = "member_deleted"

	// リクエスト
	// Header
	HEADER_CONTENT_TYPE = "application/json"
	HEADER_USER_AGENT   = "Kaonavi-Webhook"
)
