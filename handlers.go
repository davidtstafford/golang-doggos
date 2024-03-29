package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	mo "github.com/davidtstafford/golang-doggos/models"
	repo "github.com/davidtstafford/golang-doggos/repositories"
)

func DoggoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	dbClient, _ := repo.NewClient()

	doggoList, _ := dbClient.GetDoggos()

	if err := json.NewEncoder(w).Encode(doggoList); err != nil {
		panic(err)
	}
}

func AddDoggo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	dbClient, _ := repo.NewClient()
	doggo := &mo.Doggo{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, doggo); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	err = dbClient.WriteDoggo(doggo)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteDoggo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	dbClient, _ := repo.NewClient()
	doggo := &mo.Doggo{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, doggo); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	err = dbClient.DeleteDoggo(doggo)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

}
