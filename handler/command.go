package handler

import (
	"github.com/snekROmonoro/snowflake"

	"github.com/snekROmonoro/disgo/discord"
	"github.com/snekROmonoro/disgo/events"
	"github.com/snekROmonoro/disgo/rest"
)

type CommandEvent struct {
	*events.ApplicationCommandInteractionCreate
	Variables map[string]string
}

func (e *CommandEvent) GetInteractionResponse(opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().GetInteractionResponse(e.ApplicationID(), e.Token(), opts...)
}

func (e *CommandEvent) UpdateInteractionResponse(messageUpdate discord.MessageUpdate, opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().UpdateInteractionResponse(e.ApplicationID(), e.Token(), messageUpdate, opts...)
}

func (e *CommandEvent) DeleteInteractionResponse(opts ...rest.RequestOpt) error {
	return e.Client().Rest().DeleteInteractionResponse(e.ApplicationID(), e.Token(), opts...)
}

func (e *CommandEvent) GetFollowupMessage(messageID snowflake.ID, opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().GetFollowupMessage(e.ApplicationID(), e.Token(), messageID, opts...)
}

func (e *CommandEvent) CreateFollowupMessage(messageCreate discord.MessageCreate, opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().CreateFollowupMessage(e.ApplicationID(), e.Token(), messageCreate, opts...)
}

func (e *CommandEvent) UpdateFollowupMessage(messageID snowflake.ID, messageUpdate discord.MessageUpdate, opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().UpdateFollowupMessage(e.ApplicationID(), e.Token(), messageID, messageUpdate, opts...)
}

func (e *CommandEvent) DeleteFollowupMessage(messageID snowflake.ID, opts ...rest.RequestOpt) error {
	return e.Client().Rest().DeleteFollowupMessage(e.ApplicationID(), e.Token(), messageID, opts...)
}
