package gapo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Client struct {
	Host     string
	BotID    string
	BotToken string
	ApiKey   string
}

func (client Client) GetUserID(identifierCode string) (uint64, error) {
	agent := fiber.Get(client.Host + "/open-api/v2.0/exchange")
	agent.Set("Content-Type", "application/json")
	agent.Set("x-gapo-openapi-key", client.ApiKey)
	agent.QueryString("identifier_code=" + identifierCode)

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return 0, errors.Join(errs...)
	}
	if statusCode != 200 {
		return 0, fmt.Errorf("failed to get user id: %s", body)
	}

	var resp struct {
		Message string `json:"message"`
		Data    struct {
			UserId uint64 `json:"user_id"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return 0, err
	}

	return resp.Data.UserId, nil
}

func (client Client) SendMentionNotification(identifierCode string, redirectURL string) error {
	userID, err := client.GetUserID(identifierCode)
	if err != nil {
		return err
	}

	payload := map[string]any{
		"receiver_id": userID,
		"bot_id":      client.BotID,
		"body": map[string]any{
			"type":             "text",
			"text":             fmt.Sprintf("Bạn đã được đề cập trong bài viết %s", redirectURL),
			"is_markdown_text": true,
		},
	}
	marshal, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	agent := fiber.Post(client.Host + "/3rd-bot/v1.0/3rd/messages")
	agent.Set("Content-Type", "application/json")
	agent.Set("x-gapo-api-key", client.BotToken)
	agent.Body(marshal)

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	if statusCode != 200 {
		return fmt.Errorf("failed to send mention notification: %s", body)
	}

	var resp fiber.Map
	if err := json.Unmarshal(body, &resp); err != nil {
		return err
	}

	return nil
}
