package discord

import (
	"encoding/json"
	"log/slog"
	"os"
	"testing"

	"github.com/CyberDruga/fanbox2discord/src/models/discord"
)

func init() {

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)

	slog.SetDefault(slog.New(handler))

}

func TestSendMessage(t *testing.T) {
	content, err := os.ReadFile("./webhook_url")

	if err != nil {
		t.Error("[webhook_url] file is not present")
		return
	}

	response, err := SendWebhoook(SendWebhookParams{
		WebhookUrl: string(content),
		WebhookMessage: discord.WebhookMessage{
			Content:  "test",
			Username: "CyberDruga",
			AvataUrl: "https://images-ext-1.discordapp.net/external/WvUW_ugEZK8rQ13FeE2Yk_6WAa95OnupY32VeiUfNu8/%3Fsize%3D1024/https/cdn.discordapp.com/avatars/463192808468119562/ea9435d4642127dae835a2f8edd29631.webp?format=png&width=903&height=903",
		},
	})

	if err != nil {
		t.Error(err.Error())
	}

	result, err := json.Marshal(response)

	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("Result: %v", string(result))

}
