package repositories

import (
	"context"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/slog"
)

type VideoEnhanceRepository interface {
	Create(video *models.VideoEnhance) error
	FindByRequestId(requestId string) (*models.VideoEnhance, error)
	FindByUserId(userId string) ([]models.VideoEnhance, error)
	Update(response *models.VideoEnhanceResponse) error
	Delete(requestId string) error
	VideoEnhanceRepositorySetup
}

type VideoEnhanceRepositorySetup interface {
	MakeUserIdIndex()
	MakeRequestIdIndex()
}

type videoRepository struct {
	collection *mongo.Collection
}

func NewVideoEnhanceRepository(collection *mongo.Collection) VideoEnhanceRepository {
	return &videoRepository{collection}
}

func (repository *videoRepository) Create(video *models.VideoEnhance) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	inserted, err := repository.collection.InsertOne(ctx, video)
	if err != nil {
		slog.Error("Error inserting video", "err", err)
		return err
	}

	slog.Debug("Inserted video", "insertionId", inserted.InsertedID)
	return nil

}

func (repository *videoRepository) FindByRequestId(requestId string) (*models.VideoEnhance, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var video models.VideoEnhance
	err := repository.collection.FindOne(ctx, models.VideoEnhance{RequestId: requestId}).Decode(&video)
	if err != nil {
		slog.Error("Error finding video", "requestId", requestId)
		return nil, err
	}

	slog.Debug("Found video", "requestId", video.RequestId)
	return &video, nil

}

func (repository *videoRepository) FindByUserId(userId string) ([]models.VideoEnhance, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repository.collection.Find(ctx, models.VideoEnhance{UserId: userId})
	if err != nil {
		slog.Error("Error finding videos of user", "userId", userId)
		return nil, err
	}

	var videos []models.VideoEnhance
	err = cursor.All(ctx, &videos)
	if err != nil {
		return nil, err
	}

	slog.Debug("Found videos of user", "userId", userId)
	return videos, nil

}

func (repository *videoRepository) Update(response *models.VideoEnhanceResponse) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repository.collection.UpdateOne(
		ctx,
		models.VideoEnhance{RequestId: response.RequestId},
		bson.D{{Key: "$set", Value: models.VideoEnhance{EnhancedVideoUri: response.EnhancedVideoUri}}},
	)

	if err != nil {
		slog.Error("Error updating video", "requestId", response.RequestId)
		return err
	}

	slog.Debug("Updated video", "requestId", response.RequestId)
	return nil

}

func (repository *videoRepository) Delete(requestId string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repository.collection.DeleteOne(ctx, models.VideoEnhance{RequestId: requestId})
	if err != nil {
		slog.Error("Error deleting video", "requestId", requestId)
		return err
	}

	slog.Debug("Deleted video", "requestId", requestId)
	return nil

}

func (repository *videoRepository) MakeRequestIdIndex() { // used in one time setup

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexName, err := repository.collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: "requestId", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	if err != nil {
		slog.Error("Error creating requestId index", "indexName", indexName)
		panic(err)
	}

	slog.Debug("Created requestId index", "indexName", indexName)

}

func (r *videoRepository) MakeUserIdIndex() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexName, err := r.collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: "userId", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	if err != nil {
		slog.Error("Error creating userId index", "indexName", indexName)
		panic(err)
	}

	slog.Debug("Created userId index", "indexName", indexName)

}
