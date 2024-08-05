package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var Store = sessions.NewCookieStore([]byte("1122#P"))

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func HashedPasswrd(pwrd string) string {
	Hsdpwrd, err := bcrypt.GenerateFromPassword([]byte(pwrd), 14)

	if err != nil {
		fmt.Println("Unable to hash password")
		log.Fatal()
	}

	return string(Hsdpwrd)
}

// LoginHandler handles user login.
// @Summary Login
// @Description Log in to the application
// @Accept  json
// @Produce  json
// @Param   username  body  string  true  "Username"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /login [post]
func Authenticate(w http.ResponseWriter, r *http.Request) {
	var v User
	w.Header().Set("Content-Type", "application/json")
	errr := json.NewDecoder(r.Body).Decode(&v)
	if errr != nil {
		fmt.Println("Invalid")
		json.NewEncoder(w).Encode("Invalid Username or Password")
		return
	}
	data := Db.QueryRow(`SELECT password FROM USERS WHERE id=(?)`, v.Id)
	var value string
	e := data.Scan(&value)
	if e != nil {
		fmt.Println("error in Scaning", e)
	}
	er := bcrypt.CompareHashAndPassword([]byte(value), []byte(v.Password))

	// if err != nil {
	// 	fmt.Println("Error in Query")
	// }

	if er != nil {
		fmt.Println("Unable to decrypt")
		json.NewEncoder(w).Encode("Wrong Password")
	} else {
		cookie := new(http.Cookie)
		cookie.Name = "Login"
		cookie.Value = "1122#P"
		cookie.Expires = time.Date(2024, time.July, 18, 8, 8, 8, 8, time.Local)
		_, err := Db.Exec(`INSERT INTO session (userid, token, expire) VALUES (?, ?, ?);`, v.Id, cookie.Value, cookie.Expires)

		if err != nil {
			fmt.Println("Query Error", err)
		} else {
			json.NewEncoder(w).Encode("Cookie created sucessfully")
		}

		// session, er := Store.Get(r, "login")
		// if er != nil {
		// 	fmt.Println("err:", er)
		// }
		// session.Values["User"] = v.Name
		// session.Save(r, w)
		json.NewEncoder(w).Encode("Success")
	}
}

// LogoutHandler handles user logout.
// @Summary Logout
// @Description Log out from the application
// @Success 200 {object} map[string]interface{}
// @Router /logout [post]
func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var l User
	json.NewDecoder(r.Body).Decode(&l)

	data := Db.QueryRow(`select userid, token, expire from session where userid=?`, l.Id)
	var (
		userid int
		token  string
		expire []uint8
	)
	err := data.Scan(&userid, &token, &expire)
	if err != nil {
		fmt.Println(" not found", err)
		return
	}

	fmt.Println(expire)
	dateTimeString := string(expire)
	const layout = "2006-01-02 15:04:05"
	dateTime, err := time.Parse(layout, dateTimeString)
	fmt.Println(dateTime)
	fmt.Println(time.Now())
	if err != nil {
		fmt.Println("error in converting date and time ")
	}
	if userid == l.Id && time.Now().Before(dateTime) {
		res, err := Db.Exec(`update session set expire=? where userid=? `, time.Now(), l.Id)
		fmt.Println("time.Now()", time.Now().UTC())
		if err != nil {
			fmt.Println("Query error ", err)
		}
		fmt.Println(res.RowsAffected())

		json.NewEncoder(w).Encode("Your are logedout!")
	} else {
		json.NewEncoder(w).Encode("Unable yo logout")
	}
}

// SignupHandler handles user signup.
// @Summary Signup
// @Description Create a new user account
// @Accept  json
// @Produce  json
// @Param   username  body  string  true  "Username"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /signup [post]
func Register(w http.ResponseWriter, r *http.Request) {
	var u User
	//var name, password string
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		fmt.Println("database error")
	}

	_, er := Db.Exec(`INSERT INTO users (id, username, password) VALUES (?, ?, ?);`, u.Id, u.Name, HashedPasswrd(u.Password))

	if er != nil {
		fmt.Println("Query error", er)
	}

	// if er != nil {
	// 	fmt.Println("ERROR WHILE INSERTING")
	// 	log.Fatal(er)

	// } else {
	// 	fmt.Println("Record Inserted")
	// }

	// if err != nil {
	// 	fmt.Println("ERROR WHILE handling")
	// 	panic(err)
	// } else {
	// 	fmt.Println(u)

	// }

}
