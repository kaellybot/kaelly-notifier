package feeds

import (
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) GetFeedSources() ([]entities.FeedSource, error) {
	var feedSources []entities.FeedSource
	response := repo.db.GetDB().
		Model(&entities.FeedSource{}).
		Find(&feedSources)
	return feedSources, response.Error
}
