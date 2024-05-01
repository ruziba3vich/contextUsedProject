package handlers

import (
	"context"
	"contextUsedProject/internal/models"
	"contextUsedProject/internal/storage"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Comment(c *gin.Context, db *sql.DB, ctx context.Context) {
	var crdto CommentRequestDTO
	c.ShouldBindJSON(&crdto)

	twit, err := getPostById(GetPostByIdRequest{
		db:     db,
		postId: crdto.twitId,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	comment, err := storage.Comment(models.CommentDTO{
		TwitterId: crdto.twitterId,
		Twiit:     *twit,
		Comment: models.Comment{
			PostId:  twit.Id,
			Content: crdto.commentContent,
		},
		DB:      db,
		Context: context.Background(),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	// twitter, err := getTwitterById(GetTwitterByIdRequest{
	// 	db: db,
	// 	twitterId: crdto.twitterId,
	// })

	cmt, err := storage.Comment(models.CommentDTO{
		TwitterId: crdto.twitterId,
		Twiit:     *twit,
		Comment:   *comment,
		DB:        db,
		Context:   ctx,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, cmt)
}

func getPostById(gpbir GetPostByIdRequest) (*models.Twit, error) {
	query := "SELECT id, title, content FROM Twits WHERE id = $1;"
	var newTwit models.Twit
	row := gpbir.db.QueryRow(query, gpbir.postId)
	if err := row.Scan(&newTwit.Id, &newTwit.Title, &newTwit.Content); err != nil {
		return nil, err
	}
	return &newTwit, nil
}

// func getTwitterById(gpbir GetTwitterByIdRequest) (*models.Twitter, error) {
// 	query := "SELECT id, username, password FROM Twitters WHERE id = $1;"
// 	var newTwiter models.Twitter
// 	var twitterId int
// 	row := gpbir.db.QueryRow(query, gpbir.twitterId)
// 	if err := row.Scan(&twitterId, &newTwiter.Username, &newTwiter.Password); err != nil {
// 		return nil, err
// 	}
// 	newTwiter.SetId(twitterId)
// 	return &newTwiter, nil
// }

// type GetTwitterByIdRequest struct {
// 	db        *sql.DB
// 	twitterId int
// }

type CommentRequestDTO struct {
	twitterId      int
	twitId         int
	commentContent string
}

type GetPostByIdRequest struct {
	db     *sql.DB
	postId int
	// ctx context.Context
}
