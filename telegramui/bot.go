package telegramui

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type TelegramAPI struct {
	token   string
	baseUrl string
	client  http.Client
}

type Chat struct {
	ID   int    `json:"id"`
	Type string `json:"type"` // "private", "group", "supergroup" or "channel"
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type Message struct {
	ID   int    `json:"message_id"`
	From User   `json:"from"`
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Update struct {
	ID      int     `json:"update_id"`
	Message Message `json:"message"`
}

type Result interface{}

type Response[T any] struct {
	Ok     bool `json:"ok"`
	Result T    `json:"result"`
}

func NewBot(token string) *TelegramAPI {
	bot := TelegramAPI{
		baseUrl: "https://api.telegram.org/bot",
		token:   token,
		client:  http.Client{Timeout: 10 * time.Second},
	}

	return &bot
}

func (t *TelegramAPI) GetUpdates() []Update {
	methodUrl := t.baseUrl + t.token + "/getUpdates"

	resp, err := t.client.Get(methodUrl)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	buffer, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	// parse response
	response := Response[[]Update]{}

	if err = json.Unmarshal(buffer, &response); err != nil {
		return nil
	}

	updates := response.Result

	return updates
}
