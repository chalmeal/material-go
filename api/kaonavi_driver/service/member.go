package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"kaonavi_driver/components"
	"kaonavi_driver/config"
	"net/http"
	"os"
	"strconv"
)

type Member struct {
	Code string `json:"code"`
}

type Request struct {
	Event     string   `json:"event"`
	EventTime string   `json:"event_time"`
	Member    []Member `json:"member_data"`
}

// Member構造体の初期化
func NewMember() Member {
	return Member{}
}

// MemberServiceインターフェース
type MemberService interface {
	EventWebhook(event string) error
}

// Webhookイベントを処理
func (m *Member) EventWebhook(event string) error {
	// 環境変数からWebhookのURLを取得
	url := m.getWebhookURL(event)

	// リクエストボディを作成
	r := &Request{}
	members := m.createMemberCodeList()
	r.Event = event
	r.EventTime = components.GetCurrentTime()
	r.Member = members
	reqBody := r.requestToJSON(*r)

	// リクエスト生成
	req, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return err
	}

	// ヘッダーを作成
	req.Header.Set("Kaonavi-Token", os.Getenv("KAONAVI_TOKEN"))
	req.Header.Set("Content-Type", config.HEADER_CONTENT_TYPE)
	req.Header.Set("User-Agent", config.HEADER_USER_AGENT)

	// リクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer resp.Body.Close()

	// ステータスコードが200以外の場合はpanic
	if resp.StatusCode != http.StatusOK {
		panic(resp.StatusCode)
	}

	return nil
}

// WebhookのURLを取得する
func (m *Member) getWebhookURL(event string) string {
	url := os.Getenv("TEST_MODULE_API_URL")
	if event == "member_created" {
		url += os.Getenv("TEST_MODULE_CREATED_URI")
	} else if event == "member_updated" {
		url += os.Getenv("TEST_MODULE_UPDATED_URI")
	} else if event == "member_deleted" {
		url += os.Getenv("TEST_MODULE_DELETED_URI")
	} else {
		return ""
	}

	return url
}

// メンバーコードを取得する
func (m *Member) createMemberCodeList() []Member {
	// 環境変数からメンバーコードを取得
	var members []Member
	for i := 1; ; i++ {
		code := os.Getenv("MEMBER_CODE_" + strconv.Itoa(i))
		if code == "" {
			break
		}

		members = append(members, Member{Code: code})
	}

	return members
}

// Request構造体をJSON形式に変換する
func (r *Request) requestToJSON(req Request) io.Reader {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	return bytes.NewBuffer(jsonData)
}
