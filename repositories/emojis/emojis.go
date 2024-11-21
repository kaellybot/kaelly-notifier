package emojis

import (
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
	"github.com/spf13/viper"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) GetEmojis() ([]entities.Emoji, error) {
	var emojis []entities.Emoji
	response := repo.db.GetDB().
		Model(&entities.Emoji{})

	if !viper.GetBool(constants.Production) {
		response = response.
			Select("id, snowflake_dev AS snowflake, type, name, discord_name")
	}

	emojiTypes := []constants.EmojiType{
		constants.EmojiTypeItem,
		constants.EmojiTypeMisc,
	}
	response = response.
		Where("type IN (?)", emojiTypes).
		Find(&emojis)
	return emojis, response.Error
}
