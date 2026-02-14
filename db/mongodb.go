package db

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// maskMongoURI 脱敏 URI 中的密码，便于日志打印
func maskMongoURI(uri string) string {
	if uri == "" {
		return "(empty)"
	}
	// mongodb://user:password@host -> mongodb://user:****@host
	if idx := strings.Index(uri, "@"); idx > 0 {
		prefix := uri[:strings.Index(uri, "://")+3]
		rest := uri[len(prefix):]
		if at := strings.Index(rest, "@"); at > 0 {
			userPart := rest[:at]
			if colon := strings.LastIndex(userPart, ":"); colon >= 0 {
				userPart = userPart[:colon+1] + "****"
			}
			return prefix + userPart + rest[at:]
		}
	}
	return uri
}

var (
	Client   *mongo.Client
	Database *mongo.Database
)

// InitMongoDB initializes MongoDB connection
func InitMongoDB() error {
	// Get MongoDB URI from environment variable or use default
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
		log.Printf("[MongoDB] MONGODB_URI not set, using default: %s", maskMongoURI(mongoURI))
	} else {
		log.Printf("[MongoDB] MONGODB_URI from env: %s", maskMongoURI(mongoURI))
	}

	// Get database name from environment variable or use default
	dbName := os.Getenv("MONGODB_DATABASE")
	if dbName == "" {
		dbName = "ysgame"
		log.Printf("[MongoDB] MONGODB_DATABASE not set, using default: %s", dbName)
	} else {
		log.Printf("[MongoDB] MONGODB_DATABASE from env: %s", dbName)
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Ping the database
	err = Client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	Database = Client.Database(dbName)
	log.Printf("Connected to MongoDB database: %s", dbName)

	return nil
}

// CloseMongoDB closes MongoDB connection
func CloseMongoDB() error {
	if Client == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return Client.Disconnect(ctx)
}

// GetCollection returns a MongoDB collection
func GetCollection(name string) *mongo.Collection {
	return Database.Collection(name)
}
