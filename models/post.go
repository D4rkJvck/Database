package models

import (
	"database/sql"
	"time"
)

type Post struct {
	ID           string
	Author       *User
	Title        string
	Content      string
	Comments     []*Comment
	CreationDate time.Time
}

/////////////////
//	METHODS	//
////////////////

func (p *Post) Create(*sql.DB) {

}

func (p *Post) Read(*sql.DB) {

}

func (p *Post) Update(*sql.DB) {

}

func (p *Post) Delete(*sql.DB) {

}
