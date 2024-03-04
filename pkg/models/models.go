package models

import "errors"

var ErrNoRecord = errors.New("models: no matching record found")

type Item struct {
	Id       int
	Name     string
	TypeItem string
	Price    int
	ImgUrl   string
	Quantity int
}
type User struct {
	Id             int
	Email          string
	FullName       string
	HashedPassword string
	CardNumber     int
}
