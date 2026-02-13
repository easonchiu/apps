package controllers

import (
	"context"
	"net/http"
	"time"
	"ysgame/db"
	"ysgame/models"

	"github.com/gin-gonic/gin"
)

// SubmitFeedback handles feedback submission
func SubmitFeedback(c *gin.Context) {
	var req models.FeedbackRequest

	// Bind and validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Create feedback document
	feedback := models.Feedback{
		Email:     req.Email,
		Content:   req.Content,
		CreatedAt: time.Now(),
	}

	// Insert into MongoDB
	collection := db.GetCollection("feedbacks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, feedback)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save feedback"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Feedback submitted successfully",
		"id":      result.InsertedID,
	})
}
