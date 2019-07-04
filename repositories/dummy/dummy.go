package dummy

import (
	"github.com/davidtstafford/golang-doggos/models"
)

type dummyRepo struct {
	client string
}

var (
	doggo     models.Doggo
	doggoList models.Doggos
)

func NewClient() (*dummyRepo, error) {

	//doggo = models.Doggo{}
	//doggoList = make(models.Doggos, 0)

	return &dummyRepo{client: `Nope`}, nil
}

func (repo *dummyRepo) GetDoggos() (*models.Doggos, error) {

	return &doggoList, nil
}

func (repo *dummyRepo) WriteDoggo(doggo *models.Doggo) error {

	doggoList = append(doggoList, *doggo)

	return nil
}

func (repo *dummyRepo) DeleteDoggo(doggo *models.Doggo) error {

	for i, n := range doggoList {
		if n.ID == doggo.ID {
			doggoList = append(doggoList[:i], doggoList[i+1:]...)
			i--
		}

	}
	return nil
}
