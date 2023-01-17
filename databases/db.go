package databases

import (
	"database/sql"
)

var DB *sql.DB

type users struct {
	Username string
	Password string
}

func GetUsers() ([]users, error) {
	xusers := []users{}
	queryStatment, err := DB.Prepare("SELECT username, password FROM userbase;")
	if err != nil {
		return xusers, err
	}

	defer queryStatment.Close()

	rows, err := queryStatment.Query()
	if err != nil {
		return xusers, nil
	}

	for rows.Next() {
		user := users{}

		err := rows.Scan(&user.Username, &user.Password)
		if err != nil {
			return xusers, nil
		}

		xusers = append(xusers, user)
	}

	return xusers, nil

}

func InputUser(us, ps, em string) error {
	query, err := DB.Prepare(`		
	INSERT INTO userbase VALUES(?, ?, ?)
	`)

	if err != nil {
		return err
	}

	defer query.Close()

	_, err = query.Exec(us, ps, em)
	if err != nil {
		return err
	}
	return nil

}
