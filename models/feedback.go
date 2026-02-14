package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Feedback represents user feedback data
type Feedback struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email     string             `bson:"email" json:"email"`
	Content   string             `bson:"content" json:"content"`
	App       string             `bson:"app" json:"app"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

// FeedbackRequest represents the request body for feedback submission
type FeedbackRequest struct {
	Email   string `json:"email" binding:"required,email"`
	Content string `json:"content" binding:"required"`
	App     string `json:"app"`
}
