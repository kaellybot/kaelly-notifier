package emojis

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

type Repository interface {
	GetEmojis(game amqp.Game) ([]entities.Emoji, error)
}

type Impl struct {
	db databases.MySQLConnection
}
