package emojis

import (
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

type Repository interface {
	GetEmojis() ([]entities.Emoji, error)
}

type Impl struct {
	db databases.MySQLConnection
}
