package models

func Main() *Doggos {
	return &Doggos{}
}

type Doggo struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Breed string `json:"Breed"`
}

type Doggos []Doggo
