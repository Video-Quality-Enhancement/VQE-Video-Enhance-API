package repositories

import (
	"context"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/slog"
)

type VideoEnhanceRepository interface {
	Create(video *models.VideoEnhance) error
	FindByRequestId(userId, requestId string) (*models.VideoEnhance, error)
	FindAllByUserId(userId string) ([]models.VideoEnhance, error)
	Delete(userId, requestId string) error
}

type VideoEnhanceRepositorySetup interface {
	MakeUserIdIndex()
	MakeRequestIdUniqueIndex()
}

type videoEnhanceRepository struct {
	collection *mongo.Collection
}

func NewVideoEnhanceRepository(collection *mongo.Collection) VideoEnhanceRepository {
	return &videoEnhanceRepository{collection}
}

func NewVideoEnhanceRepositorySetup(collection *mongo.Collection) VideoEnhanceRepositorySetup {
	return &videoEnhanceRepository{collection}
}

func (repository *videoEnhanceRepository) Create(video *models.VideoEnhance) error {

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

func (repository *videoEnhanceRepository) FindByRequestId(userId, requestId string) (*models.VideoEnhance, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "requestId", Value: requestId}, {Key: "userId", Value: userId}}

	var video models.VideoEnhance
	err := repository.collection.FindOne(ctx, filter).Decode(&video)
	// * added userId along with requestId coz only that user should be able to the video with that particular requestId

	if err != nil {
		slog.Error("Error finding video", "requestId", requestId)
		return nil, err
	}

	slog.Debug("Found video", "requestId", video.RequestId)
	return &video, nil

}

func (repository *videoEnhanceRepository) FindAllByUserId(userId string) ([]models.VideoEnhance, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "userId", Value: userId}}

	cursor, err := repository.collection.Find(ctx, filter)
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

func (repository *videoEnhanceRepository) Delete(userId, requestId string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "requestId", Value: requestId}, {Key: "userId", Value: userId}}

	deleteResult, err := repository.collection.DeleteOne(ctx, filter)

	if err != nil {
		slog.Error("Error deleting video", "requestId", requestId, "err", err)
		return err
	}

	slog.Debug("Deleted video", "requestId", requestId, "deleteCount", deleteResult.DeletedCount)
	return nil

}

func (repository *videoEnhanceRepository) MakeRequestIdUniqueIndex() { // used in one time setup

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

func (r *videoEnhanceRepository) MakeUserIdIndex() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexName, err := r.collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: "userId", Value: 1}},
			Options: options.Index(),
		},
	)

	if err != nil {
		slog.Error("Error creating userId index", "indexName", indexName)
		panic(err)
	}

	slog.Debug("Created userId index", "indexName", indexName)

}
