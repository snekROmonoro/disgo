package handlers

import (
	"slices"

	"github.com/snekROmonoro/disgo/bot"
	"github.com/snekROmonoro/disgo/cache"
	"github.com/snekROmonoro/disgo/discord"
	"github.com/snekROmonoro/disgo/events"
	"github.com/snekROmonoro/disgo/gateway"
	"github.com/snekROmonoro/snowflake"
)

type updatedEmoji struct {
	old discord.Emoji
	new discord.Emoji
}

func gatewayHandlerGuildEmojisUpdate(client bot.Client, sequenceNumber int, shardID int, event gateway.EventGuildEmojisUpdate) {
	client.EventManager().DispatchEvent(&events.EmojisUpdate{
		GenericEvent:           events.NewGenericEvent(client, sequenceNumber, shardID),
		EventGuildEmojisUpdate: event,
	})

	if client.Caches().CacheFlags().Missing(cache.FlagEmojis) {
		return
	}

	createdEmojis := map[snowflake.ID]discord.Emoji{}
	deletedEmojis := map[snowflake.ID]discord.Emoji{}
	updatedEmojis := map[snowflake.ID]updatedEmoji{}

	client.Caches().EmojisForEach(event.GuildID, func(emoji discord.Emoji) {
		deletedEmojis[emoji.ID] = emoji
	})

	for _, newEmoji := range event.Emojis {
		oldEmoji, ok := deletedEmojis[newEmoji.ID]
		if ok {
			delete(deletedEmojis, newEmoji.ID)
			if isEmojiUpdated(oldEmoji, newEmoji) {
				updatedEmojis[newEmoji.ID] = updatedEmoji{new: newEmoji, old: oldEmoji}
			}
			continue
		}
		createdEmojis[newEmoji.ID] = newEmoji
	}

	for _, emoji := range createdEmojis {
		client.Caches().AddEmoji(emoji)
		client.EventManager().DispatchEvent(&events.EmojiCreate{
			GenericEmoji: &events.GenericEmoji{
				GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
				GuildID:      event.GuildID,
				Emoji:        emoji,
			},
		})
	}

	for _, emoji := range updatedEmojis {
		client.Caches().AddEmoji(emoji.new)
		client.EventManager().DispatchEvent(&events.EmojiUpdate{
			GenericEmoji: &events.GenericEmoji{
				GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
				GuildID:      event.GuildID,
				Emoji:        emoji.new,
			},
			OldEmoji: emoji.old,
		})
	}

	for _, emoji := range deletedEmojis {
		client.Caches().RemoveEmoji(event.GuildID, emoji.ID)
		client.EventManager().DispatchEvent(&events.EmojiDelete{
			GenericEmoji: &events.GenericEmoji{
				GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
				GuildID:      event.GuildID,
				Emoji:        emoji,
			},
		})
	}

}

func isEmojiUpdated(old discord.Emoji, new discord.Emoji) bool {
	if old.Name != new.Name {
		return true
	}
	if !slices.Equal(old.Roles, new.Roles) {
		return true
	}
	return false
}
