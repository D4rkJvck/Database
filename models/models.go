package models

import (
	"net/http"
	"time"
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
	Session *http.Cookie
}

type Post struct {
	ID string
	Author *User
	Title string
	Content string
	Comments []*Comment
	CreationDate time.Time
}

type Comment struct {
	ID string
	Author *User
	Target *Post
	Content string
	CreationDate time.Time
}

type Message struct {
	ID string
	Sender *User
	Receiver *User
	Content string
	CreationDate time.Time
	Seen bool
}
