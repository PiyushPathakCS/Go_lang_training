package usecase

import (
	"clear_code/database"
)

type Val struct {
	Id       int
	Name     string
	Password string
}
type Srv struct {
	dbconnect database.UserDatabase
}
type Servce interface {
	Signup(data *Val) error
	Login(id int)
}

func NewSrv(dbconnect database.UserDatabase) *Srv {
	return &Srv{
		dbconnect: dbconnect,
	}
}

func (srv *Srv) Signup(data *Val) error {
	var db_data database.User
	db_data.Id = data.Id
	db_data.Name = data.Name
	db_data.Password = data.Password
	err := srv.dbconnect.InsertUser(&db_data)

	if err != nil {
		return err
	}
	return nil
}

func (srv *Srv) Login(id int) {

	err := srv.dbconnect.GetUser(id)

	if err != nil {
		return
	}

}
