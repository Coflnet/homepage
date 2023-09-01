package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func SendContactMessage(firstname, lastname, email, message string) error {
	if wrongString(firstname) {
		return errors.New("Firstname is empty")
	}

	if wrongString(lastname) {
		return errors.New("Lastname is empty")
	}

	if wrongString(email) {
		return errors.New("Email is empty")
	}

	if wrongString(message) {
		return errors.New("Message is empty")
	}

	// call discord webhook url
	url, err := loadDiscordWebhookUrl()
	if err != nil {
		return err
	}

	return sendDiscordMessage(url, messageToSend(firstname, lastname, email, message))
}

func wrongString(str string) bool {
	if str == "" {
		return true
	}

	if len(str) > 1000 {
		return true
	}

	return false
}

func messageToSend(firstname, lastname, email, message string) string {
	return fmt.Sprintf("Firstname: %s\nLastname: %s\nEmail: %s\nMessage: %s", firstname, lastname, email, message)
}

func loadDiscordWebhookUrl() (string, error) {
	v := os.Getenv("DISCORD_WEBHOOK_URL")
	if v == "" {
		return "", errors.New("discord webhook url is not set")
	}
	return v, nil
}

func sendDiscordMessage(url, message string) error {
	data := &WebhookRequest{
		Content:             message,
		Username:            "homepage-forward",
		AvatarUrl:           "https://cdn.discordapp.com/app-icons/888725077191974913/0c0e3b97e6865091ef14162083a54a42.png?size=256",
		AllowedMentionsData: AllowedMentions{Parse: make([]string, 0)},
	}

	body, err := json.Marshal(data)
	if err != nil {
		slog.Error("error when marshalling webhook request", err)
		return err
	}

	_, err = http.DefaultClient.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		slog.Error("error when sending webhook request", err)
		return err
	}

	return nil
}

type WebhookRequest struct {
	Content             string          `json:"content"`
	Username            string          `json:"username"`
	AvatarUrl           string          `json:"avatar_url"`
	AllowedMentionsData AllowedMentions `json:"allowed_mentions"`
}

type AllowedMentions struct {
	Parse []string `json:"parse"`
}
