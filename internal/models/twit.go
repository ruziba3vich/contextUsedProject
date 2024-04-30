package models

type Twit struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	userId  int
}

func (twit *Twit) GetComment(cdto CommentDTO) (*Comment, error) {
	query := `
		INSERT INTO Comments(twitter_id, post_id, content)
			VALUES ($1, $2, $3) RETURNING id, content;
	`
	var newComment Comment
	row := cdto.DB.QueryRow(query, twit.userId, twit.Id, cdto.Comment.Content)
	if err := row.Scan(&newComment.Id, &newComment.Content); err != nil {
		return nil, err
	}

	return &newComment, nil
}
