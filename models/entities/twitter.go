package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

// TwitterAccount is identified by (id, game): each game has its own accounts and its
// own news channel, so a lookup by ID alone could return another game's channel.
type TwitterAccount struct {
	ID            string    `gorm:"primaryKey"`
	Game          amqp.Game `gorm:"primaryKey"`
	Name          string
	NewsChannelID string
}
