package emojis

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	repository "github.com/kaellybot/kaelly-notifier/repositories/emojis"
)

type Service interface {
	GetMiscStringEmoji(emojiID constants.EmojiMiscID) string
	GetItemTypeStringEmoji(itemType amqp.ItemType) string
}

type Impl struct {
	emojiStore map[constants.EmojiType]emojiStore
	repository repository.Repository
}

type emojiStore map[string]entities.Emoji
