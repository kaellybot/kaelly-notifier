package emojis

import (
	"fmt"
	"strings"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	repository "github.com/kaellybot/kaelly-notifier/repositories/emojis"
	"github.com/rs/zerolog/log"
)

func New(repository repository.Repository) (*Impl, error) {
	emojis, err := repository.GetEmojis()
	if err != nil {
		return nil, err
	}

	log.Info().
		Int(constants.LogEntityCount, len(emojis)).
		Msgf("Emojis loaded")

	emojiStore := make(map[constants.EmojiType]emojiStore)
	for _, emoji := range emojis {
		innerStore, found := emojiStore[emoji.Type]
		if !found {
			innerStore = make(map[string]entities.Emoji)
			emojiStore[emoji.Type] = innerStore
		}

		innerStore[emoji.ID] = emoji
	}

	return &Impl{
		emojiStore: emojiStore,
		repository: repository,
	}, nil
}

func (service *Impl) GetMiscStringEmoji(emojiMiscID constants.EmojiMiscID) string {
	innerStore, found := service.emojiStore[constants.EmojiTypeMisc]
	if !found {
		log.Warn().
			Str(constants.LogEmojiType, string(constants.EmojiTypeMisc)).
			Msgf("No miscellaneous type store found, returning empty emoji")
		return mapEmojiString(nil)
	}

	emojiID := string(emojiMiscID)
	emoji, found := innerStore[emojiID]
	if !found {
		log.Warn().
			Str(constants.LogEntityID, emojiID).
			Msgf("No miscellaneous emoji found, returning empty emoji")
		return mapEmojiString(nil)
	}

	return mapEmojiString(&emoji)
}

func (service *Impl) GetItemTypeStringEmoji(itemType amqp.ItemType) string {
	innerStore, found := service.emojiStore[constants.EmojiTypeItem]
	if !found {
		log.Warn().
			Str(constants.LogEmojiType, string(constants.EmojiTypeItem)).
			Msgf("No item type store found, returning empty emoji")
		return mapEmojiString(nil)
	}

	emojiID := itemType.String()
	emoji, found := innerStore[emojiID]
	if !found {
		log.Warn().
			Str(constants.LogEntityID, emojiID).
			Msgf("No item type emoji found, returning empty emoji")
		return mapEmojiString(nil)
	}

	return mapEmojiString(&emoji)
}

func mapEmojiString(emoji *entities.Emoji) string {
	if emoji != nil {
		if len(strings.TrimSpace(emoji.ID)) > 0 {
			return fmt.Sprintf("<:%v:%v>", emoji.DiscordName, emoji.Snowflake)
		}

		return emoji.Name
	}

	return ""
}
