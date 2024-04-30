package handlers

import (
	"contextUsedProject/internal/models"
	"contextUsedProject/internal/storage"
)

func Comment() {
	storage.Comment(models.CommentDTO{
		
	})
}