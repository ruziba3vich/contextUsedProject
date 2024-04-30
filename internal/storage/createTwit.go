package storage

import (
	"context"
	"contextUsedProject/internal/models"
	"errors"
	"time"
)

type ResponsePostChannel struct {
	twit models.Twit
	err  error
}

func CreateTwit(cdto models.CreateTwitDTO) (*models.Twit, error) {
	ctx, cancelFunc := context.WithTimeout(cdto.Context, time.Millisecond*100)
	defer cancelFunc()

	responseChan := make(chan ResponsePostChannel)

	go func() {
		twit, err := cdto.Twitter.Twit(
			models.TwitDTO{
				Twit: cdto.Twit,
				DB:   cdto.DB,
			})
		responseChan <- ResponsePostChannel{
			twit: *twit,
			err:  err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("server took too long to execute")
		case newChan := <-responseChan:
			return &newChan.twit, newChan.err
		}
	}
}
