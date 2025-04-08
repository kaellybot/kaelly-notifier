package news

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/repositories/almanaxes"
	"github.com/kaellybot/kaelly-notifier/repositories/feeds"
	"github.com/kaellybot/kaelly-notifier/repositories/twitter"
)

func New(almanaxRepo almanaxes.Repository, feedRepo feeds.Repository,
	twitterRepo twitter.Repository) (*Impl, error) {
	almanaxNews, errAlm := almanaxRepo.GetAlmanaxNews()
	if errAlm != nil {
		return nil, errAlm
	}

	feedSources, errFeed := feedRepo.GetFeedSources()
	if errFeed != nil {
		return nil, errFeed
	}

	twitterAccounts, errTwitter := twitterRepo.GetTwitterAccounts()
	if errTwitter != nil {
		return nil, errTwitter
	}

	return &Impl{
		almanaxNews:     almanaxNews,
		feedSources:     feedSources,
		twitterAccounts: twitterAccounts,
		almanaxRepo:     almanaxRepo,
		feedRepo:        feedRepo,
		twitterRepo:     twitterRepo,
	}, nil
}

func (service *Impl) GetAlmanaxNews(locale amqp.Language,
	game amqp.Game) *entities.AlmanaxNews {
	for _, almanax := range service.almanaxNews {
		if almanax.Locale == locale &&
			almanax.Game == game {
			return &almanax
		}
	}
	return nil
}

func (service *Impl) GetFeedSource(feedTypeID string, locale amqp.Language,
	game amqp.Game) *entities.FeedSource {
	for _, feedSource := range service.feedSources {
		if feedSource.FeedTypeID == feedTypeID &&
			feedSource.Locale == locale &&
			feedSource.Game == game {
			return &feedSource
		}
	}
	return nil
}

func (service *Impl) GetTwitterAccount(accountID string) *entities.TwitterAccount {
	for _, account := range service.twitterAccounts {
		if account.ID == accountID {
			return &account
		}
	}
	return nil
}
