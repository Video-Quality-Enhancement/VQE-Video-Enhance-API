package repositories

import (
	"context"
	"log"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/account/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccountRepository interface {
	UpsertAccount(account *models.Account) error
	FindAccount(email string) (*models.Account, error)
	AccountRepositorySetup
}

type AccountRepositorySetup interface {
	MakeEmailIndex()
}

type accountRepository struct {
	collection *mongo.Collection
}

func NewAccountRepository(collection *mongo.Collection) AccountRepository {
	return &accountRepository{collection}
}

func (repository *accountRepository) UpsertAccount(account *models.Account) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := models.Account{Email: account.Email}
	update := bson.D{{Key: "$set", Value: account}}
	opts := options.Update().SetUpsert(true)

	result, err := repository.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Println("Error upserting account to database ", account)
		return err
	}

	log.Println("Upserted account with email: ", account.Email)
	log.Printf("Is account updated: %v\n", result.ModifiedCount)
	log.Printf("Is account upserted: %v\n", result.UpsertedCount)
	return nil

}

func (repository *accountRepository) FindAccount(email string) (*models.Account, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var account models.Account
	err := repository.collection.FindOne(ctx, models.Account{Email: email}).Decode(&account)
	if err != nil {
		log.Println("Account not found in DB: ", err)
		return nil, err
	}

	log.Println("Got account from DB: ", account)
	return &account, nil

}

func (repository *accountRepository) MakeEmailIndex() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexName, err := repository.collection.Indexes().CreateOne(
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
