package repositories

import (
	"github.com/davidtstafford/golang-doggos/models"
)

type DBClient interface {
	GetDoggos() (*models.Doggos, error)
	WriteDoggo(doggo *models.Doggo) error
	DeleteDoggo(doggo *models.Doggo) error
}
