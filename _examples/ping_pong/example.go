package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/snekROmonoro/disgo"
	"github.com/snekROmonoro/disgo/bot"
	"github.com/snekROmonoro/disgo/discord"
	"github.com/snekROmonoro/disgo/events"
	"github.com/snekROmonoro/disgo/gateway"
)

func main() {
	slog.Info("starting example...")
	slog.Info("disgo version", slog.String("version", disgo.Version))

	client, err := disgo.New(os.Getenv("disgo_token"),
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuildMessages,
				gateway.IntentMessageContent,
			),
		),
		bot.WithEventListenerFunc(onMessageCreate),
	)
	if err != nil {
		slog.Error("error while building disgo", slog.String("err", err.Error()))
		return
	}

	defer client.Close(context.TODO())

	if err = client.OpenGateway(context.TODO()); err != nil {
		slog.Error("errors while connecting to gateway", slog.String("err", err.Error()))
		return
	}

	slog.Info("example is now running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}

func onMessageCreate(event *events.MessageCreate) {
	if event.Message.Author.Bot {
		return
	}
	var message string
	if event.Message.Content == "ping" {
		message = "pong"
	} else if event.Message.Content == "pong" {
		message = "ping"
	}
	if message != "" {
		_, _ = event.Client().Rest().CreateMessage(event.ChannelID, discord.NewMessageCreateBuilder().SetContent(message).Build())
	}
}
