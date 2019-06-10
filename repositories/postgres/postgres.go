package postgres

import (
	"database/sql"
	"fmt"
	"os"

	// Used in conjunction with database/sql" to provide Postgres driver
	_ "github.com/lib/pq"

	"github.com/davidtstafford/golang-doggos/models"
	repo "github.com/davidtstafford/golang-doggos/repositories"
)

var (
	host     string
	port     string
	user     string
	password string
	dbname   string
)

type postgresRepo struct {
	client *sql.DB
}

func NewClient() (repo.DBClient, error) {
	loadOSEnvs()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &postgresRepo{client: db}, nil
}

func loadOSEnvs() {
	host = os.Getenv("pgHost")
	port = os.Getenv("pgPort")
	user = os.Getenv("pgUser")
	password = os.Getenv("pgPassword")
	dbname = os.Getenv("pgDbName")
}

func (repo *postgresRepo) GetDoggos() (*models.Doggos, error) {
	doggo := models.Doggo{}
	doggoList := make(models.Doggos, 0)

	rows, err := repo.client.Query(`select "ID", "Name", "Breed" from demo.doggos`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&doggo.ID, &doggo.Name, &doggo.Breed)
		if err != nil {
			return nil, err
		}
		doggoList = append(doggoList, doggo)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	repo.client.Close()

	return &doggoList, nil
}

func (repo *postgresRepo) WriteDoggo(doggo *models.Doggo) error {

	sqlStatement := `INSERT INTO demo.doggos ("ID", "Name", "Breed" ) VALUES ($1, $2, $3)`
	_, err := repo.client.Exec(sqlStatement, &doggo.ID, &doggo.Name, &doggo.Breed)
	if err != nil {
		return err
	}

	return nil
}

func (repo *postgresRepo) DeleteDoggo(doggo *models.Doggo) error {

	sqlStatement := `DELETE FROM demo.doggos WHERE "ID" = $1`
	_, err := repo.client.Exec(sqlStatement, &doggo.ID)
	if err != nil {
		return err
	}

	return nil
}
