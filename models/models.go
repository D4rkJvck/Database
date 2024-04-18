package models

type User struct {
	ID         string 
	Nickname   string
	Age        int
	Gender     string
	First_Name string
	Last_Name  string
	Email      string
	Password   []byte
}
