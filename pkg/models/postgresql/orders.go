package postgresql

import (
	"KazakhAliexpress/SE2220/pkg/models"
	"database/sql"
	"log"
)

type OrderModel struct {
	DB *sql.DB
}

func (m *OrderModel) Insert(UserId int, address, msg string) {
	stmt := "INSERT INTO orders(user_id,address,msg) values($1,$2,$3)"
	_, err := m.DB.Exec(stmt, UserId, address, msg)
	if err != nil {
		log.Fatal(err)
	}
}
func (m *OrderModel) GetOrdersByUserId(UserId int) (*models.Order, error) {
	stmt := "SELECT * FROM orders where user_id=$1"
	row := m.DB.QueryRow(stmt, UserId)
	order := &models.Order{}
	err := row.Scan(&order.Id, &order.UserId, &order.Address, &order.Msg)
	if err != nil {
		log.Println(err)
	}

	return order, nil
}
