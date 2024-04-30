package handlers

import (
	"contextUsedProject/internal/storage"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadTwits(c *gin.Context, db *sql.DB) {
	twits, err := storage.GetTwits(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, twits)
}
