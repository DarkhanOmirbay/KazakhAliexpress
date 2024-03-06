package postgresql

import (
	"KazakhAliexpress/SE2220/pkg/models"
	"database/sql"
	"errors"
	"log"
)

type CartModel struct {
	DB *sql.DB
}

func (m *CartModel) AddCart(ItemId, UserId int) {
	stmt := "INSERT INTO carts(item_id,user_id) values ($1,$2)"
	_, err := m.DB.Exec(stmt, ItemId, UserId)
	if err != nil {
		log.Println(err)
	}
}
func (m *CartModel) GetCartsById(id int) ([]*models.Cart, error) {
	stmt := "SELECT * from carts where user_id=$1"
	rows, err := m.DB.Query(stmt, id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	carts := []*models.Cart{}
	for rows.Next() {
		cart := &models.Cart{}
		err := rows.Scan(&cart.Id, &cart.ItemId, &cart.UserId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, models.ErrNoRecord
			} else {
				return nil, err
			}
		}
		carts = append(carts, cart)
	}
	return carts, nil
}
func (m *CartModel) DeleteItem(ItemId int) {
	stmt := "DELETE FROM carts where item_id=$1"
	_, err := m.DB.Exec(stmt, ItemId)
	if err != nil {
		log.Println(err)
	}
}
func (m *CartModel) DeleteItemUseUserId(UserId int) {
	stmt := "DELETE FROM carts where user_id=$1"
	_, err := m.DB.Exec(stmt, UserId)
	if err != nil {
		log.Println(err)
	}
}
