package postgresql

import (
	"KazakhAliexpress/SE2220/pkg/models"
	"database/sql"
	"errors"
	"log"
)

type ItemModel struct {
	DB *sql.DB
}

func (m *ItemModel) Insert(name, type_item, imgurl string, price, quantity int) {
	stmt := "INSERT INTO items(name, type_item, price, img_url, quantity) values($1,$2,$3,$4,$5)"
	res, err := m.DB.Exec(stmt, name, type_item, price, imgurl, quantity)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
func (m *ItemModel) Read() ([]*models.Item, error) {
	stmt := "SELECT * FROM items"
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*models.Item{}
	for rows.Next() {
		item := &models.Item{}
		err := rows.Scan(&item.Id, &item.Name, &item.TypeItem, &item.Price, &item.ImgUrl, &item.Quantity)
		if err != nil {
			return nil, err
		}
		items = append(items, item)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
func (m *ItemModel) GetItem(id int) (*models.Item, error) {
	stmt := "SELECT * from items WHERE id=$1"

	row := m.DB.QueryRow(stmt, id)
	item := &models.Item{}
	err := row.Scan(&item.Id, &item.Name, &item.TypeItem, &item.Price, &item.ImgUrl, &item.Quantity)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return item, nil
}
func (m *ItemModel) Update(name, type_item, imgurl string, id, price, quantity int) error {
	stmt := "UPDATE items SET name =$1, type_item = $2, price = $3, img_url = $4, quantity = $5 WHERE id = $6"
	_, err := m.DB.Exec(stmt, name, type_item, price, imgurl, quantity, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
func (m *ItemModel) Delete(id int) error {
	stmt := "DELETE FROM items where id=$1"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
