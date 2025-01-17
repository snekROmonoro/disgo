package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"
	"github.com/snekROmonoro/disgo"
	"github.com/snekROmonoro/disgo/bot"
	"github.com/snekROmonoro/disgo/discord"
	"github.com/snekROmonoro/disgo/events"
	"github.com/snekROmonoro/disgo/httpserver"
	"github.com/snekROmonoro/snowflake"
)

var (
	token     = os.Getenv("disgo_token")
	publicKey = os.Getenv("disgo_public_key")
	guildID   = snowflake.GetEnv("disgo_guild_id")

	commands = []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			Name:        "say",
			Description: "says what you say",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionString{
					Name:        "message",
					Description: "What to say",
					Required:    true,
				},
				discord.ApplicationCommandOptionBool{
					Name:        "ephemeral",
					Description: "If the response should only be visible to you",
					Required:    true,
				},
			},
		},
	}
)

func main() {
	slog.Info("starting example...")
	slog.Info("disgo version", slog.String("version", disgo.Version))

	// use custom ed25519 verify implementation
	httpserver.Verify = func(publicKey httpserver.PublicKey, message, sig []byte) bool {
		return ed25519.Verify(publicKey, message, sig)
	}

	client, err := disgo.New(token,
		bot.WithHTTPServerConfigOpts(publicKey,
			httpserver.WithURL("/interactions/callback"),
			httpserver.WithAddress(":80"),
		),
		bot.WithEventListenerFunc(commandListener),
	)
	if err != nil {
		panic("error while building disgo instance: " + err.Error())
	}

	defer client.Close(context.TODO())

	if _, err = client.Rest().SetGuildCommands(client.ApplicationID(), guildID, commands); err != nil {
		panic("error while registering commands: " + err.Error())
	}

	if err = client.OpenHTTPServer(); err != nil {
		panic("error while starting http server: " + err.Error())
	}

	slog.Info("example is now running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}

func commandListener(event *events.ApplicationCommandInteractionCreate) {
	data := event.SlashCommandInteractionData()
	if data.CommandName() == "say" {
		if err := event.CreateMessage(discord.NewMessageCreateBuilder().
			SetContent(data.String("message")).
			SetEphemeral(data.Bool("ephemeral")).
			Build(),
		); err != nil {
			event.Client().Logger().Error("error on sending response: ", err)
		}
	}
}
