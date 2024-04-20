package models

import (
	"database/sql"
	"time"
)

type Message struct {
	ID           string
	Sender       *User
	Receiver     *User
	Content      string
	CreationDate time.Time
	Seen         bool
}

/////////////////
//	METHODS	//
////////////////

func (msg *Message) Create(*sql.DB) {
	
}

func (msg *Message) Read(*sql.DB) {

}

func (msg *Message) Update(*sql.DB) {

}

func (msg *Message) Delete(*sql.DB) {

}
