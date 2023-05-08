package repositories

import (
	"context"
	"log"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VideoEnhanceRepository interface {
	Create(*models.VideoEnhance) error
	FindByRequestId(string) (*models.VideoEnhance, error)
	FindByEmail(email string) ([]models.VideoEnhance, error)
	Update(response *models.VideoEnhanceResponse) error
	Delete(string) error
	VideoEnhanceRepositorySetup
}

type VideoEnhanceRepositorySetup interface {
	MakeEmailIndex()
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
		log.Println("Error Inserting video to database ", video)
		return err
	}

	log.Println("Inserted video with id: ", inserted.InsertedID)
	return nil

}

func (repository *videoRepository) FindByRequestId(requestId string) (*models.VideoEnhance, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var video models.VideoEnhance
	err := repository.collection.FindOne(ctx, models.VideoEnhance{RequestId: requestId}).Decode(&video)
	if err != nil {
		log.Println("Error finding video with id: ", requestId)
		return nil, err
	}

	log.Println("Found video with id: ", video.RequestId)
	return &video, nil

}

func (repository *videoRepository) FindByEmail(email string) ([]models.VideoEnhance, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repository.collection.Find(ctx, models.VideoEnhance{Email: email})
	if err != nil {
		log.Println("Error finding videos with email: ", email)
		return nil, err
	}

	var videos []models.VideoEnhance
	err = cursor.All(ctx, &videos)
	if err != nil {
		return nil, err
	}

	log.Println("Found videos with email: ", email)
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
		log.Println("Error updating video with id: ", response.RequestId)
		return err
	}

	log.Println("Updated video with id: ", response.RequestId)
	return nil

}

func (repository *videoRepository) Delete(requestId string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repository.collection.DeleteOne(ctx, models.VideoEnhance{RequestId: requestId})
	if err != nil {
		log.Println("Error deleting video with id: ", requestId)
		return err
	}

	log.Println("Deleted video with id: ", requestId)
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
		log.Println("Error creating index with name: ", indexName)
		panic(err)
	}

	log.Println("Created index with name: ", indexName)

}

func (r *videoRepository) MakeEmailIndex() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexName, err := r.collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	if err != nil {
		log.Println("Error creating index with name: ", indexName)
		panic(err)
	}

	log.Println("Created index with name: ", indexName)

}
