package repositories

import (
	"context"
	"errors"

	"github.com/thyms-c/be-memo-app/internal/configs"
	"github.com/thyms-c/be-memo-app/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CounterRepository interface {
	GetByName(ctx context.Context, name string) (*models.Counter, error)
	Create(ctx context.Context, name string) (*models.Counter, error)
	Increment(ctx context.Context, name string) error
}

type counterRepositoryImpl struct {
	collection *mongo.Collection
}

func NewCounterRepository(configs *configs.Config, mongoClient *mongo.Client) CounterRepository {
	return &counterRepositoryImpl{
		collection: mongoClient.Database(configs.MongoDatabase).Collection("counters"),
	}
}

// Create implements CounterRepository.
func (c *counterRepositoryImpl) Create(ctx context.Context, name string) (*models.Counter, error) {

	// Create a new counter with the initial value of 0
	counter := models.Counter{
		Name:  name,
		Value: 0,
	}

	_, err := c.collection.InsertOne(ctx, counter)
	if err != nil {
		return nil, err
	}

	return &counter, nil
}

// GetByName implements CounterRepository.
func (c *counterRepositoryImpl) GetByName(ctx context.Context, name string) (*models.Counter, error) {
	var counter *models.Counter

	err := c.collection.FindOne(ctx, bson.M{"name": name}).Decode(&counter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return counter, nil
}

// Increment implements CounterRepository.
func (c *counterRepositoryImpl) Increment(ctx context.Context, name string) error {
	update := bson.M{
		"$inc": bson.M{"value": 1},
	}

	result, err := c.collection.UpdateOne(ctx, bson.M{"name": name}, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
