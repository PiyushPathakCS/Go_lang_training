package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type UserDatabase interface {
	InsertUser(u *User) error
	GetUser(id int) error
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type db struct {
	conn *sql.DB
}

func Connect() *db {
	var data = db{}
	conn, err := sql.Open("mysql", "root:Piyush@99@tcp(localhost:3306)/user")
	if err != nil {
		log.Fatal(err)
	} else {
		err = conn.Ping()
		if err != nil {

			log.Fatal(err)
		}
		fmt.Println("Connected to database!")
	}

	data.conn = conn

	return &data
}

func (i *db) InsertUser(u *User) error {
	_, err := i.conn.Exec(`INSERT INTO users (id,username,password) VALUES (?,?,?)`, u.Id, u.Name, u.Password)

	if err != nil {
		fmt.Println("Error during Signup!", err)
		return err
	} else {
		fmt.Println("Success you have signed up ")
	}

	return nil

}

func (i *db) GetUser(id int) error {
	data := i.conn.QueryRow(`select id,name,password from users where id = ?`, id)
	var u User

	err := data.Scan(&u.Id, &u.Name, &u.Password)

	if err != nil {
		fmt.Println("Unable to get user info ", err)
		return err
	}

	return nil
}

func (i *db) UpdateUser(u *User) {
	_, err := i.conn.Exec(`update users set name=?, password=? where id=?`, u.Name, u.Password, u.Id)

	if err != nil {
		fmt.Println("Unable to update your info", err)
	}
}

func (i *db) DeleteUser(id int) {
	_, err := i.conn.Exec(`delete from users where id=?`, id)

	if err != nil {
		fmt.Println("Unable to delete user", err)
	}
}
