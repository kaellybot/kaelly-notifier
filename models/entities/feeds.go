package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

type FeedSource struct {
	Locale        amqp.Language `gorm:"primaryKey"`
	Game          amqp.Game     `gorm:"primaryKey"`
	FeedTypeID    string        `gorm:"primaryKey"`
	NewsChannelID string
}
