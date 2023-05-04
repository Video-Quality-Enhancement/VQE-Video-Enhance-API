package repositories

import (
	"context"
	"log"

	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/interfaces"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VideoRepository interface {
	interfaces.Repository
	Create(*models.Video) error
	FindByCorrelationId(string) (*models.Video, error)
	FindByEmail(email string) ([]models.Video, error)
	Update(string, string) error
	Delete(string) error
	MakeCorrelationIdIndex()
}

type videoRepository struct {
	collection *mongo.Collection
}

func NewVideoRepository(collection *mongo.Collection) VideoRepository {
	return &videoRepository{collection}
}

func (repository *videoRepository) Create(video *models.Video) error {

	inserted, err := repository.collection.InsertOne(context.Background(), video)
	if err != nil {
		return err
	}

	log.Println("Inserted video with id: ", inserted.InsertedID)
	return nil

}

func (repository *videoRepository) FindByCorrelationId(correlationId string) (*models.Video, error) {

	var video models.Video
	err := repository.collection.FindOne(context.Background(), models.Video{CorrelationId: correlationId}).Decode(&video)
	if err != nil {
		return nil, err
	}

	log.Println("Found video with id: ", video.CorrelationId)
	return &video, nil

}

func (repository *videoRepository) FindByEmail(email string) ([]models.Video, error) {

	cursor, err := repository.collection.Find(context.Background(), models.Video{Email: email})
	if err != nil {
		return nil, err
	}

	var videos []models.Video
	err = cursor.All(context.Background(), &videos)
	if err != nil {
		return nil, err
	}

	log.Println("Found videos with email: ", email)
	return videos, nil

}

func (repository *videoRepository) Update(correlationId string, enhancedVideoUrl string) error {

	_, err := repository.collection.UpdateOne(
		context.Background(),
		models.Video{CorrelationId: correlationId},
		bson.D{{Key: "$set", Value: models.Video{EnhancedVideoUrl: enhancedVideoUrl}}},
	)

	if err != nil {
		return err
	}

	log.Println("Updated video with id: ", correlationId)
	return nil

}

func (repository *videoRepository) Delete(correlationId string) error {

	_, err := repository.collection.DeleteOne(context.Background(), models.Video{CorrelationId: correlationId})
	if err != nil {
		return err
	}

	log.Println("Deleted video with id: ", correlationId)
	return nil

}

func (repository *videoRepository) MakeCorrelationIdIndex() { // used in one time setup

	indexName, err := repository.collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "correlationId", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	if err != nil {
		panic(err)
	}

	log.Println("Created index with name: ", indexName)

}
