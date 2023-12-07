package models

type Medicine struct {
	Id     string `json: id`
	Name   string `json: name`
	Docmed bool   `json: docmed`
}
