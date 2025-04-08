package almanaxes

import (
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

type Repository interface {
	GetAlmanaxNews() ([]entities.AlmanaxNews, error)
}

type Impl struct {
	db databases.MySQLConnection
}
