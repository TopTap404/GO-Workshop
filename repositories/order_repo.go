
package repositories

import (
	"workshop/database"
	"workshop/models"
)

func CreateOrder(o *models.Order) error {
	return database.DB.Create(o).Error
}
