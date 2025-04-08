package almanaxes

import (
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) GetAlmanaxNews() ([]entities.AlmanaxNews, error) {
	var almanaxNews []entities.AlmanaxNews
	response := repo.db.GetDB().
		Model(&entities.AlmanaxNews{}).
		Find(&almanaxNews)
	return almanaxNews, response.Error
}
