package main

import (
	"kaonavi_driver/service"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	member service.Member
)

func init() {
	// MemberサービスDI
	member = service.NewMember()

	// 環境変数の読み込み
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	event := os.Args[1]
	if event == "" {
		log.Fatal("event is required")
	}

	if err := member.EventWebhook(event); err != nil {
		log.Fatal(err)
	}
}
