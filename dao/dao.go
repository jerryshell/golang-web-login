package dao

import (
	"database/sql"
	"log"

	"jerryshell.cn/login_demo/domain"
)

func getDB() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./user.db")
	checkError(err)
	return
}

// FindUserByUsernameAndPassword 通过 username 和 password 查找 User
func FindUserByUsernameAndPassword(username string, password string) (user *domain.User) {
	sql := "select id, email from user where username=? and password=?"

	db := getDB()
	defer db.Close()

	rows, err := db.Query(sql, username, password)
	checkError(err)
	defer rows.Close()

	if rows.Next() {
		var id int
		var email string
		rows.Scan(&id, &email)

		user = &domain.User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
	}
	return
}

// AddUser 添加新 User
func AddUser(user *domain.User) {
	sql := "insert into user(username, password, email) values(?,?,?)"

	db := getDB()
	defer db.Close()

	_, err := db.Exec(sql, user.Username, user.Password, user.Email)
	checkError(err)
}

// UpdateUser 更新 User
func UpdateUser(user *domain.User) {
	sql := "update user set username=?, password=?, email=? where id=?"

	db := getDB()
	defer db.Close()

	_, err := db.Exec(sql, user.Username, user.Password, user.Email, user.ID)
	checkError(err)
}

// DeleteUser 删除 User
func DeleteUser(id int) {
	sql := "delete from user where id=?"

	db := getDB()
	defer db.Close()

	_, err := db.Exec(sql, id)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}