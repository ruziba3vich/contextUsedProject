package models

type Twitter struct {
	id       int
	Username string `json:"username"`
	Password string `json:"password"`
}

func (twitter *Twitter) SetId(id int) {
	twitter.id = id
}

func (twitter *Twitter) SetUsername(username string) {
	twitter.Username = username
}

func (twitter *Twitter) SetPassword(pwd string) {
	twitter.Password = pwd
}

func (twitter *Twitter) Twit(twdto TwitDTO) (*Twit, error) {
	query := `
		INSERT INTO Twits(user_id, title, content)
			VALUES ($1, $2, $3, $4)
			RETURNING title, content;
	`
	row := twdto.DB.QueryRow(query, twitter.id, twdto.Twit.Title, twdto.Twit.Content)
	var newTwit Twit

	if err := row.Scan(&newTwit.Title, &newTwit.Content); err != nil {
		return nil, err
	}
	return &newTwit, nil
}

func (twitter *Twitter) Comment(cdto *CommentDTO) (*Comment, error) {
	cdto.TwitterId = twitter.id
	*cdto.Comment.twitterId = twitter.id
	return cdto.Twiit.GetComment(*cdto)
}
