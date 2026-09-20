package emojis

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
	"github.com/spf13/viper"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

// GetEmojis loads the snowflakes of one game's Discord application. The notifier
// serves every game from a single process, so it loads once per game and renders
// each message with its own game's emojis.
//
// The join is a LEFT one on purpose: an emoji with no snowflake for this
// application keeps an empty one, and the caller falls back on its unicode name
// rather than rendering another application's snowflake.
func (repo *Impl) GetEmojis(game amqp.Game) ([]entities.Emoji, error) {
	var emojis []entities.Emoji
	emojiTypes := []constants.EmojiType{
		constants.EmojiTypeItem,
		constants.EmojiTypeMisc,
	}

	response := repo.db.GetDB().
		Model(&entities.Emoji{}).
		Select("emojis.id, emojis.type, emojis.name, emojis.discord_name, "+
			"COALESCE(emoji_snowflakes.snowflake, '') AS snowflake").
		Joins("LEFT JOIN emoji_snowflakes "+
			"ON emoji_snowflakes.emoji_id = emojis.id "+
			"AND emoji_snowflakes.emoji_type = emojis.type "+
			"AND emoji_snowflakes.game = ? AND emoji_snowflakes.production = ?",
			game, viper.GetBool(constants.Production)).
		Where("emojis.type IN (?)", emojiTypes).
		Find(&emojis)
	return emojis, response.Error
}
