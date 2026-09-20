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
	store := make(map[amqp.Game]map[constants.EmojiType]emojiStore)
	for _, game := range constants.GetServedGames() {
		emojis, err := repository.GetEmojis(game.AMQPGame)
		if err != nil {
			return nil, err
		}

		gameStore := make(map[constants.EmojiType]emojiStore)
		missing := 0
		for _, emoji := range emojis {
			innerStore, found := gameStore[emoji.Type]
			if !found {
				innerStore = make(map[string]entities.Emoji)
				gameStore[emoji.Type] = innerStore
			}

			innerStore[emoji.ID] = emoji
			if emoji.Snowflake == "" {
				missing++
			}
		}

		store[game.AMQPGame] = gameStore
		log.Info().
			Str(constants.LogGame, game.AMQPGame.String()).
			Int(constants.LogEntityCount, len(emojis)).
			Msgf("Emojis loaded")

		// Reported once, at load: these render as their unicode name instead.
		if missing > 0 {
			log.Warn().
				Str(constants.LogGame, game.AMQPGame.String()).
				Int(constants.LogEntityCount, missing).
				Msgf("Emojis without a snowflake for %v, falling back on their name",
					game.BotName)
		}
	}

	return &Impl{
		emojiStore: store,
		repository: repository,
	}, nil
}

func (service *Impl) GetMiscStringEmoji(game amqp.Game,
	emojiMiscID constants.EmojiMiscID) string {
	return service.get(game, constants.EmojiTypeMisc, string(emojiMiscID))
}

func (service *Impl) GetItemTypeStringEmoji(game amqp.Game, itemType amqp.ItemType) string {
	return service.get(game, constants.EmojiTypeItem, itemType.String())
}

func (service *Impl) get(game amqp.Game, emojiType constants.EmojiType,
	emojiID string) string {
	gameStore, found := service.emojiStore[game]
	if !found {
		log.Warn().
			Str(constants.LogGame, game.String()).
			Msgf("No emoji store found for this game, returning empty emoji")
		return mapEmojiString(nil)
	}

	innerStore, found := gameStore[emojiType]
	if !found {
		log.Warn().
			Str(constants.LogGame, game.String()).
			Str(constants.LogEmojiType, string(emojiType)).
			Msgf("No emoji type store found, returning empty emoji")
		return mapEmojiString(nil)
	}

	emoji, found := innerStore[emojiID]
	if !found {
		log.Warn().
			Str(constants.LogGame, game.String()).
			Str(constants.LogEntityID, emojiID).
			Msgf("No emoji found, returning empty emoji")
		return mapEmojiString(nil)
	}

	return mapEmojiString(&emoji)
}

// mapEmojiString renders the snowflake, not the ID: an emoji row always has an ID,
// but only has a snowflake for the applications it was uploaded to. Rendering
// without one would produce a broken reference, so the unicode name is used.
func mapEmojiString(emoji *entities.Emoji) string {
	if emoji == nil {
		return ""
	}

	if len(strings.TrimSpace(emoji.Snowflake)) > 0 {
		return fmt.Sprintf("<:%v:%v>", emoji.DiscordName, emoji.Snowflake)
	}

	return emoji.Name
}
