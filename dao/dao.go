package dao

import (
	"database/sql"
	"log"

	"github.com/jerryshell/golang-web-login/domain"
)

const createUserTable = "CREATE TABLE `user` ( `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, `username` TEXT NOT NULL, `password` TEXT NOT NULL, `email` TEXT );"

const createAdmin = "INSERT INTO user(id, username, password, email) VALUES(1, 'admin', 'admin', 'admin@admin.com');"

func init() {
	initEnv()
	initAdmin()
}

func initEnv() {
	db, err := sql.Open("sqlite3", "./user.db")
	checkError(err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	_, err = db.Exec(createUserTable)
	if err != nil {
		log.Println(err)
	}
}

func initAdmin() {
	db, err := sql.Open("sqlite3", "./user.db")
	checkError(err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	_, err = db.Exec(createAdmin)
	if err != nil {
		log.Println(err)
	}
}

func getDB() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./user.db")
	checkError(err)
	return
}

// FindUserByUsernameAndPassword 通过 username 和 password 查找 User
func FindUserByUsernameAndPassword(username string, password string) *domain.User {
	sqlStr := "select id, email from user where username=? and password=?"

	db := getDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	rows, err := db.Query(sqlStr, username, password)
	checkError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	if rows.Next() {
		var id int
		var email string
		err := rows.Scan(&id, &email)
		if err != nil {
			log.Println(err)
			return nil
		}

		return &domain.User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
	}
	return nil
}

// AddUser 添加新 User
func AddUser(user *domain.User) {
	sqlStr := "insert into user(username, password, email) values(?,?,?)"

	db := getDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	_, err := db.Exec(sqlStr, user.Username, user.Password, user.Email)
	checkError(err)
}

// UpdateUser 更新 User
func UpdateUser(user *domain.User) {
	sqlStr := "update user set username=?, password=?, email=? where id=?"

	db := getDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	_, err := db.Exec(sqlStr, user.Username, user.Password, user.Email, user.ID)
	checkError(err)
}

// DeleteUser 删除 User
func DeleteUser(id int) {
	sqlStr := "delete from user where id=?"

	db := getDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	_, err := db.Exec(sqlStr, id)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
