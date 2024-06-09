package external_hook

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

type WorkflowHook struct {
	host        string
	webhookUuid string
}

func NewWorkflowHookClient(host string, webhookUuid string) *WorkflowHook {
	return &WorkflowHook{host: host, webhookUuid: webhookUuid}
}

func (client WorkflowHook) SendTopic(ctx context.Context, req TopicReq) error {
	marshal, err := json.Marshal(req)
	if err != nil {
		return err
	}

	agent := fiber.Post(client.host + "/webhook/" + client.webhookUuid)
	agent.Set("Content-Type", "application/json")
	agent.Body(marshal)

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	if statusCode != 200 {
		return fmt.Errorf("failed sent hook workflow: %s", body)
	}

	var resp fiber.Map
	if err := json.Unmarshal(body, &resp); err != nil {
		return err
	}

	log.Println("sent hook workflow with req:", string(marshal))

	return nil
}
