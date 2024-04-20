package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID           string
	Author       *User
	Target       *Post
	Content      string
	CreationDate time.Time
}

/////////////////
//	METHODS	//
////////////////

func (c *Comment) Create(*sql.DB) {

}

func (c *Comment) Read(*sql.DB) {

}

func (c *Comment) Update(*sql.DB) {

}

func (c *Comment) Delete(*sql.DB) {

}
