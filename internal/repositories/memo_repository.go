package repositories

import (
	"context"

	"github.com/thyms-c/be-memo-app/internal/configs"
	"github.com/thyms-c/be-memo-app/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MemoRepository interface {
	GetAll(ctx context.Context) ([]*models.Memo, error)
	Create(ctx context.Context, memo *models.Memo) (*models.Memo, error)
	GetByUserType(ctx context.Context, userType string) ([]*models.Memo, error)
}

type memoRepositoryImpl struct {
	collection *mongo.Collection
}

func NewMemoRepository(configs *configs.Config, mongoClient *mongo.Client) MemoRepository {
	return &memoRepositoryImpl{
		collection: mongoClient.Database(configs.MongoDatabase).Collection("memos"),
	}
}

// Create implements MemoRepository.
func (m *memoRepositoryImpl) Create(ctx context.Context, memo *models.Memo) (*models.Memo, error) {
	memo.ID = primitive.NewObjectID()

	_, err := m.collection.InsertOne(ctx, memo)
	if err != nil {
		return nil, err
	}

	return memo, nil
}

// GetAll implements MemoRepository.
func (m *memoRepositoryImpl) GetAll(ctx context.Context) ([]*models.Memo, error) {
	var memos []*models.Memo

	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := m.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var memo models.Memo
		if err := cursor.Decode(&memo); err != nil {
			return nil, err
		}
		memos = append(memos, &memo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return memos, nil
}

// GetByUserType implements MemoRepository.
func (m *memoRepositoryImpl) GetByUserType(ctx context.Context, userType string) ([]*models.Memo, error) {
	var memos []*models.Memo

	sort := 1

	if userType == string(models.AdminRole) {
		sort = -1
	}

	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: sort}})

	cursor, err := m.collection.Find(ctx, bson.M{"user_type": userType}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var memo models.Memo
		if err := cursor.Decode(&memo); err != nil {
			return nil, err
		}
		memos = append(memos, &memo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return memos, nil
}
