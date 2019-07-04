package repositories

import (
	"os"

	"github.com/davidtstafford/golang-doggos/models"
	"github.com/davidtstafford/golang-doggos/repositories/dummy"
	"github.com/davidtstafford/golang-doggos/repositories/dynamo"
	"github.com/davidtstafford/golang-doggos/repositories/postgres"
)

var (
	dbType string
)

type DBClient interface {
	GetDoggos() (*models.Doggos, error)
	WriteDoggo(doggo *models.Doggo) error
	DeleteDoggo(doggo *models.Doggo) error
}

// NewClient will return the specfic client made by the underlying database repo
// and will return the one that is defined dbType variable
func NewClient() (DBClient, error) {
	loadOSEnvs()
	switch dbType {
	case "POSTGRES":
		return postgres.NewClient()
	case "DYNAMO":
		return dynamo.NewClient()
	default:
		return dummy.NewClient()
	}
	return nil, nil
}

func loadOSEnvs() {
	dbType = os.Getenv("dbType")
}
