package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type User struct {
	id int
}

func GetUserInfo(userId int) (User, error) {
	DB, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/my_db")
	if err := DB.Ping(); err != nil {
		return User{}, errors.Wrap(err, "open mysql error")
	}

	var user User
	err := DB.QueryRow(fmt.Sprintf("select id from user where id = %d", userId)).Scan(user)
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, errors.Wrap(err, fmt.Sprintf("find user null,user id: %d", userId))
	}
	return user, nil
}
