package storage

import (
	"context"
	"contextUsedProject/internal/models"
	"errors"
	"time"
)

type ResponseCommentChannel struct {
	comment models.Comment
	err     error
}

func Comment(cdto models.CommentDTO) (*models.Comment, error) {
	ctx, cancelFunc := context.WithTimeout(cdto.Context, time.Millisecond*100)
	responseChannel := make(chan ResponseCommentChannel)

	defer cancelFunc()

	go func() {
		comment, err := cdto.Twiit.GetComment(cdto)
		responseChannel <- ResponseCommentChannel{
			comment: *comment,
			err:     err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("server took too long to execute")
		case newChan := <-responseChannel:
			return &newChan.comment, nil
		}
	}
}
