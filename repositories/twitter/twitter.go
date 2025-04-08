package twitter

import (
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) GetTwitterAccounts() ([]entities.TwitterAccount, error) {
	var twitterAccounts []entities.TwitterAccount
	response := repo.db.GetDB().
		Model(&entities.TwitterAccount{}).
		Find(&twitterAccounts)
	return twitterAccounts, response.Error
}
