package models

type Comment struct {
	Id        int `json:"id"`
	twitterId *int
	PostId    int    `json:"post_id"`
	Content   string `json:"content"`
}
