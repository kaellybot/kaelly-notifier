package emojis

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	repository "github.com/kaellybot/kaelly-notifier/repositories/emojis"
)

// The getters take the game because a snowflake only renders for the application
// that owns it: this single process posts for every game, so a message must be
// built with the emojis of its own game.
type Service interface {
	GetMiscStringEmoji(game amqp.Game, emojiID constants.EmojiMiscID) string
	GetItemTypeStringEmoji(game amqp.Game, itemType amqp.ItemType) string
}

type Impl struct {
	emojiStore map[amqp.Game]map[constants.EmojiType]emojiStore
	repository repository.Repository
}

type emojiStore map[string]entities.Emoji
