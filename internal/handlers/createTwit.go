package handlers

import (
	"context"
	"contextUsedProject/internal/models"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTwit(c *gin.Context, db *sql.DB) {
	var (
		crtwr models.CreatTwitRequest
	)
	ctx := context.Background()
	c.ShouldBindJSON(&crtwr)

	twitter, err := getTwitterFromDatabase(GetTwitterFromDatabaseCall{
		crtwr.Username,
		crtwr.Password,
		db,
		ctx,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	twit, err := twitter.Twit(models.TwitDTO{
		Twit: models.Twit{
			Title:   crtwr.Title,
			Content: crtwr.Content,
		},
		DB: db,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, twit)
}

type TwitGetter struct {
	twitter *models.Twitter
	err     error
}

func getTwitterFromDatabase(gtfdc GetTwitterFromDatabaseCall) (*models.Twitter, error) {
	ctx, cancelFunc := context.WithTimeout(gtfdc.cnx, time.Millisecond*200)
	defer cancelFunc()
	chnl := make(chan TwitGetter)
	var id int
	var twitter models.Twitter
	query := "SELECT id FROM Twitters WHERE username = $1 AND password = $2;"
	go func() {
		err := gtfdc.db.QueryRow(query, gtfdc.username, gtfdc.password).Scan(&id)
		if err != nil {
			chnl <- TwitGetter{
				nil,
				err,
			}
		}
		twitter.SetId(id)
		twitter.SetUsername(gtfdc.username)
		twitter.SetPassword(gtfdc.password)
		chnl <- TwitGetter{
			&twitter,
			nil,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("time out")
		case result := <-chnl:
			return result.twitter, result.err
		}
	}
}

type GetTwitterFromDatabaseCall struct {
	username string
	password string
	db       *sql.DB
	cnx      context.Context
}
