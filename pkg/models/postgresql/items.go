package postgresql

import (
	"KazakhAliexpress/SE2220/pkg/models"
	"database/sql"
	"log"
)

type ItemModel struct {
	DB *sql.DB
}

func (m *ItemModel) Insert(db *sql.DB, name, type_item, imgurl string, price, quantity int) {
	stmt := "INSERT INTO items(name, type_item, price, img_url, quantity) values($1,$2,$3,$4,$5)"
	res, err := db.Exec(stmt, name, type_item, price, imgurl, quantity)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
func (m *ItemModel) Read(db *sql.DB) ([]*models.Item, error) {
	stmt := "SELECT * FROM items"
	rows, err := db.Query(stmt)
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
func (m *ItemModel) GetItem(db *sql.DB, id int) (*models.Item, error) {
	stmt := "SELECT * from items WHERE id=$1"

	row := db.QueryRow(stmt, id)
	item := &models.Item{}
	err := row.Scan(&item.Id, &item.Name, &item.TypeItem, &item.Price, &item.ImgUrl, &item.Quantity)
	if err != nil {
		return nil, err
	}

	return item, nil
}
