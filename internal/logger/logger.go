package logger

import (
	"STVbotGo/internal/backend"
	"fmt"
	"log"
)

var (
	url      = "https://discord.com/api/webhooks/1538275361752875038/2uVpfg-_agvG-H9UVChyFK0V9hKIrXgYy9sYX1fyuw4BpXaZ6WqxKLSqejmJcMFsBEnZ"
	username = "STVbot log"
)

func Info(logMsg string) {
	message := backend.Webhook{
		Username: &username,
		Content:  fmt.Sprintf("```\nINFO: %v\n```", logMsg),
	}
	err := backend.SendWebhook(url, message)
	if err != nil {
		log.Printf("failed to send webhook: %v", err)
	}
}

func Warn(logMsg string) {
	message := backend.Webhook{
		Username: &username,
		Content:  fmt.Sprintf("```\nWARN: %v\n```", logMsg),
	}
	err := backend.SendWebhook(url, message)
	if err != nil {
		log.Printf("failed to send webhook: %v", err)
	}
}

func Error(logMsg string) {
	message := backend.Webhook{
		Username: &username,
		Content:  fmt.Sprintf("```\nERROR: %v\n```", logMsg),
	}
	err := backend.SendWebhook(url, message)
	if err != nil {
		log.Printf("failed to send webhook: %v", err)
	}
}
