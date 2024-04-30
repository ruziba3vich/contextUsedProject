package models

import (
	"context"
	"database/sql"
)

type CommentDTO struct {
	TwitterId int
	Twiit     Twit
	Comment   Comment
	DB        *sql.DB
	Context context.Context
}

type TwitDTO struct {
	Twit Twit
	DB   *sql.DB
}

type CreateTwitDTO struct {
	Twitter Twitter
	Twit    Twit
	DB      *sql.DB
	Context context.Context
}

type CreatTwitRequest struct {
	Username string
	Password string
	Title string
	Content string
}
