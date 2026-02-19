package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"ysgame/db"
	"ysgame/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SubmitFeedback handles feedback submission
func SubmitFeedback(c *gin.Context) {
	var req models.FeedbackRequest

	// Bind and validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	app := c.Query("app")

	// Create feedback document
	feedback := models.Feedback{
		Email:     req.Email,
		Content:   req.Content,
		App:       app,
		Status:    "pending",
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

// ListFeedback returns paginated feedbacks sorted by created_at desc
func ListFeedback(c *gin.Context) {
	// Parse pagination params
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	collection := db.GetCollection("feedbacks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Count total documents
	total, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count feedbacks"})
		return
	}

	// Query with pagination and sort
	skip := int64((page - 1) * pageSize)
	limit := int64(pageSize)
	opts := options.Find().
		SetSort(bson.D{{Key: "created_at", Value: -1}}).
		SetSkip(skip).
		SetLimit(limit)

	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query feedbacks"})
		return
	}
	defer cursor.Close(ctx)

	var feedbacks []models.Feedback
	if err := cursor.All(ctx, &feedbacks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode feedbacks"})
		return
	}

	// Return empty array instead of null
	if feedbacks == nil {
		feedbacks = []models.Feedback{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":     feedbacks,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

// BatchDeleteFeedback deletes multiple feedbacks by IDs
func BatchDeleteFeedback(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	objectIDs := make([]primitive.ObjectID, 0, len(req.IDs))
	for _, id := range req.IDs {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID: " + id})
			return
		}
		objectIDs = append(objectIDs, oid)
	}

	collection := db.GetCollection("feedbacks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": objectIDs}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete feedbacks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": result.DeletedCount})
}

// BatchUpdateFeedbackStatus updates the status of multiple feedbacks
func BatchUpdateFeedbackStatus(c *gin.Context) {
	var req struct {
		IDs    []string `json:"ids" binding:"required"`
		Status string   `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if req.Status != "pending" && req.Status != "processed" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status must be 'pending' or 'processed'"})
		return
	}

	objectIDs := make([]primitive.ObjectID, 0, len(req.IDs))
	for _, id := range req.IDs {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID: " + id})
			return
		}
		objectIDs = append(objectIDs, oid)
	}

	collection := db.GetCollection("feedbacks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.UpdateMany(ctx, bson.M{"_id": bson.M{"$in": objectIDs}}, bson.M{"$set": bson.M{"status": req.Status}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update feedback status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated": result.ModifiedCount})
}
