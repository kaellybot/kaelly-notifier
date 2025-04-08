package news

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/repositories/almanaxes"
	"github.com/kaellybot/kaelly-notifier/repositories/feeds"
	"github.com/kaellybot/kaelly-notifier/repositories/twitter"
)

type Service interface {
	GetAlmanaxNews(locale amqp.Language, game amqp.Game) *entities.AlmanaxNews
	GetFeedSource(feedTypeID string, locale amqp.Language, game amqp.Game) *entities.FeedSource
	GetTwitterAccount(accountID string) *entities.TwitterAccount
}

type Impl struct {
	almanaxNews     []entities.AlmanaxNews
	feedSources     []entities.FeedSource
	twitterAccounts []entities.TwitterAccount
	almanaxRepo     almanaxes.Repository
	feedRepo        feeds.Repository
	twitterRepo     twitter.Repository
}
