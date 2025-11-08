package main

import (
	"context"
	"embed"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"time"

	"github.com/CyberDruga/fanbox2discord/db"
	"github.com/CyberDruga/fanbox2discord/src/api/discord"
	"github.com/CyberDruga/fanbox2discord/src/config"
	"github.com/CyberDruga/fanbox2discord/src/generic"
	"github.com/CyberDruga/fanbox2discord/src/json-template"
	"github.com/davecgh/go-spew/spew"
)

//go:embed db/migrations/*.sql
var fs embed.FS

func main() {

	logLevel := slog.LevelInfo

	if i := slices.Index(os.Args, "--debug"); i != -1 {
		logLevel = slog.LevelDebug
	}

	opts := &slog.HandlerOptions{
		AddSource: slices.Contains(os.Args, "--show-source"),
		Level:     logLevel,
	}

	var handler slog.Handler = slog.NewTextHandler(os.Stdout, opts)
	if slices.Contains(os.Args, "--json") {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	slog.SetDefault(slog.New(handler))

	var configFile = "./config.toml"
	slices.Index(os.Args, "--config")

	if i := slices.Index(os.Args, "--config"); i != -1 {
		configFile = os.Args[i+1]
	}

	conf, err := config.LoadConfig(configFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		return
	}

	database, err := db.NewClient(conf.Database)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(1)
	}

	populate := slices.Contains(os.Args, "--populate")

	var ctx = context.Background()

	loop := func() {

		defer slog.Info("Done!")

		for _, acc := range conf.Accounts {
			posts, err := generic.GetLatestPosts(acc.CreatorId)

			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err.Error())
				continue
			}

			slices.Reverse(posts)

			slog.Info("Checking for new posts")

			for _, post := range posts {

				slog.Info("Checking post", "postId", post.PostId, "creatorId", post.CreatorId)
				_, err := database.GetPost(ctx, post.PostId)

				// gives error if it doesn't exist
				if err == nil {
					continue
				}

				if !populate {

					slog.Info("Trying to post stuff on discord.", "postId", post.PostId, "creatorId", post.CreatorId)

					message, err := json_template.ApplyTemplate(acc.NewMessageTemplate, post)

					if err != nil {
						slog.Error(err.Error())
						continue
					}

					var webhookUrl = conf.WebhookUrl

					if acc.WebhookUrl != "" {
						webhookUrl = acc.WebhookUrl
					}

					_, err = discord.SendWebhoook(discord.SendWebhookParams{
						WebhookUrl:     webhookUrl,
						WebhookMessage: message,
					})

					if err != nil {
						slog.Error(err.Error())
						continue
					}

				}

				if populate {
					populate = false
				}

				_, err = database.SaveNewPost(ctx, db.SaveNewPostParams{
					PostID:    post.PostId,
					CreatorID: post.CreatorId,
				})

				if err != nil {
					slog.Error(err.Error())
					continue
				}

			}

		}

	}

	if slices.Contains(os.Args, "--print-example") {
		for _, acc := range conf.Accounts {
			posts, err := generic.GetLatestPosts(acc.CreatorId)
			if err != nil {
				slog.Error(err.Error())
				os.Exit(1)
			}

			spew.Dump(posts[0])
		}
		os.Exit(0)
	}

	if conf.Repeat.Enable {
		for {
			loop()
			time.Sleep(conf.Repeat.EveryXSeconds * time.Second)
		}
	} else {
		loop()
	}

}
