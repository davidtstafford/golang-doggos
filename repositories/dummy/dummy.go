package dummy

import (
	"github.com/davidtstafford/golang-doggos/models"
	repo "github.com/davidtstafford/golang-doggos/repositories"
)

type dummyRepo struct {
	client string
}

func NewClient() (repo.DBClient, error) {

	return &dummyRepo{client: `Nope`}, nil
}

func (repo *dummyRepo) GetDoggos() (*models.Doggos, error) {
	doggo := models.Doggo{
		ID:    "1",
		Name:  "Blake",
		Breed: "Long Hair Chi",
	}

	doggoList := make(models.Doggos, 0)
	doggoList = append(doggoList, doggo)

	return &doggoList, nil
}

func (repo *dummyRepo) WriteDoggo(doggo *models.Doggo) error {

	return nil
}

func (repo *dummyRepo) DeleteDoggo(doggo *models.Doggo) error {

	return nil
}
