package models

import "errors"

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Item struct {
	Id       int
	Name     string
	TypeItem string
	Price    int
	ImgUrl   string
	Quantity int
}
type User struct {
	Id               int
	Email            string
	FullName         string
	HashedPassword   []byte
	HashedCardNumber []byte
	ExpirationDate   string
	CVV              int
}

type Cart struct {
	Id     int
	UserId int
	ItemId int
}

type Order struct {
	Id      int
	UserId  int
	Address string
	Msg     string
}
