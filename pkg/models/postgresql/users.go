package postgresql

import (
	"KazakhAliexpress/SE2220/pkg/models"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(email, fullname, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users ( email,full_name, hashed_password)
VALUES($1,$2,$3)`

	_, err = m.DB.Exec(stmt, email, fullname, string(hashedPassword))
	if err != nil {
		var pgError *pq.Error
		if errors.As(err, &pgError) {
			if pgError.Code == "23505" && strings.Contains(pgError.Message, "users_email_key") {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}
	return nil

}

func (m *UserModel) Authenticate(email, password string) (int, error) {

	var id int
	var hashedPassword []byte
	stmt := "SELECT id, hashed_password FROM users WHERE email = $1 "
	row := m.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil

}

func (m *UserModel) Update(id int, cardNumber, expDate, cvv string) error {
	hashedCardNumber, err := bcrypt.GenerateFromPassword([]byte(cardNumber), 12)
	if err != nil {
		return err
	}
	stmt := "UPDATE users SET hashed_card_number=$1,expiration_date=$2,cvv=$3 WHERE id=$4"
	_, err = m.DB.Exec(stmt, string(hashedCardNumber), expDate, cvv, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
