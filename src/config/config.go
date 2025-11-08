package config

import (
	"time"

	"dario.cat/mergo"
	"github.com/BurntSushi/toml"
	"github.com/go-errors/errors"
)

type Config struct {
	Database   string    `toml:"database"`
	WebhookUrl string    `toml:"webhook-url"`
	Accounts   []Account `toml:"account"`
	Repeat     Repeat    `toml:"repeat"`
}

type Repeat struct {
	Enable        bool          `toml:"enable"`
	EveryXSeconds time.Duration `toml:"every-x-seconds"`
}

type Account struct {
	CreatorId          string `toml:"creator-id"`
	WebhookUrl         string `toml:"webhook-url"`
	NewMessageTemplate string `toml:"new-message-template"`
}

func LoadConfig(filePath string) (config Config, err error) {

	_, err = toml.DecodeFile(filePath, &config)
	if err != nil {
		err = errors.Join(errors.Errorf("Coudln't load config file"), err)
		return
	}

	mergo.Merge(&config, Config{
		Database: ":memory:",
	})

	if config.WebhookUrl == "" {
		err = errors.Errorf("[webhook-url] needs to be informed on the config file")
		return
	}

	if len(config.Accounts) == 0 {
		err = errors.Errorf("At least one [[account]] needs to be configured")
		return
	}

	for i, acc := range config.Accounts {
		if acc.CreatorId == "" {
			err = errors.Errorf("account number %d doesn't have [creator-id] set", i+1)
			return
		}

		if acc.NewMessageTemplate == "" {
			err = errors.Errorf("account number %d doesn't have [new-message-template] set", i+1)
			return
		}

	}

	return
}
