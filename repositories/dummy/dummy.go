//dummy isn't to be consider a mock test.  It's purpose is to allow quick dev against something without
// having to worry about database conectivity.  In essence by flipping an env variable one doesn't need to worry about
// db configurations.  Useful for first deploy to ensure other aspects are working
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
