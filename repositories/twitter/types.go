package twitter

import (
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

type Repository interface {
	GetTwitterAccounts() ([]entities.TwitterAccount, error)
}

type Impl struct {
	db databases.MySQLConnection
}
