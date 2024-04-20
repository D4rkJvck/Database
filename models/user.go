package models

import (
	"database/sql"
	"net/http"
)

type User struct {
	ID         string
	Nickname   string
	Age        int
	Gender     string
	First_Name string
	Last_Name  string
	Email      string
	Password   []byte
	Session    *http.Cookie
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//	METHODS																		//
//////////////////////////////////////////////////////////////////////////////////////////////////////

func (u *User) Create(db *sql.DB) error {
	// Prepare Query
	stmt, err_pre := db.Prepare(`
	INSERT INTO users (user_id, nickname, age, gender, first_name, last_name, email, password)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err_pre != nil {
		return err_pre
	}
	defer stmt.Close()
	
	// Execute Query
	_, err_exe := stmt.Exec(
		u.ID, u.Nickname,
		u.Age,
		u.Gender,
		u.First_Name,
		u.Last_Name,
		u.Email,
		u.Password,
	)
	if err_exe != nil {
		return err_exe
	}
	
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func (u *User) Read(db *sql.DB) error {
	// TODO: SELECT
	return nil
}

// TODO: func (u *User) Update(db *sql.DB, data ...interface{}) {}

// TODO: func (u *User) Delete(db *sql.DB) {}
