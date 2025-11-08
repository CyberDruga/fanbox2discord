package discord

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	webhook "github.com/CyberDruga/fanbox2discord/src/models/discord"
	"github.com/go-errors/errors"
)

type SendWebhookParams struct {
	WebhookUrl string
	webhook.WebhookMessage
}

func SendWebhoook(params SendWebhookParams) (response webhook.WebhookResponse, err error) {

	if params.WebhookUrl == "" {
		panic("[WebhookUrl] can never be empty")
	}

	body, err := json.Marshal(params.WebhookMessage)

	if err != nil {
		err = errors.Join(errors.Errorf("Couldn't turn message into json"), err)
		slog.Error(err.Error())
		return
	}

	res, err := http.Post(params.WebhookUrl+"?wait=true", "application/json", bytes.NewBuffer(body))

	if err != nil {
		err = errors.Join(errors.Errorf("Couldn't send post request"), err)
		slog.Error(err.Error())
		return
	}

	defer res.Body.Close()

	stuff, err := io.ReadAll(res.Body)

	slog.Debug("Content of response: " + string(stuff))

	err = json.Unmarshal(stuff, &response)

	if err != nil {
		err = errors.Join(errors.New("Coudln't parse response from json"), err)
		slog.Error(err.Error())
		return
	}

	if response.Message != "" {
		err = errors.New(response.Message)
		slog.Error(
			err.Error(),
			"body", string(body),
		)
		return
	}

	return
}
