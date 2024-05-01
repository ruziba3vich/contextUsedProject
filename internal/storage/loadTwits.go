package storage

import (
	"context"
	"database/sql"
	"time"
)

var page int = 20
var next = 20

func GetTwits(db *sql.DB) (twits []TwitResponse, e error) {
	query := `
		SELECT t.id, twitter.username, t.title, t.content FROM Twits t
		INNER JOIN Twitters twitter ON twitter.id = t.twitter_id
		LIMIT $1;
	`
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancelFunc()

	rows, err := db.QueryContext(ctx, query, page)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var twit TwitResponse
		err := rows.Scan(
			&twit.TwitId,
			&twit.TwitterUsername,
			&twit.TwitTitle,
			&twit.TwitContent)
		if err != nil {
			return nil, err
		}
		twits = append(twits, twit)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	for range next {
		page++
	}
	return twits, nil
}

type TwitResponse struct {
	TwitId          int    `json:"twit_id"`
	TwitterUsername string `json:"twitter_username"`
	TwitTitle       string `json:"twit_title"`
	TwitContent     string `json:"twit_content"`
}
