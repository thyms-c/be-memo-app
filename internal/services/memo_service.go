package services

import (
	"context"
	"fmt"
	"time"

	"github.com/thyms-c/be-memo-app/internal/models"
	"github.com/thyms-c/be-memo-app/internal/repositories"
	"github.com/thyms-c/be-memo-app/internal/requests"
)

type MemoService interface {
	GetAllMemos(ctx context.Context) ([]*models.Memo, error)
	CreateMemo(ctx context.Context, req *requests.MemoRequest, userRole models.Role) (*models.Memo, error)
	GetMemoByUserType(ctx context.Context, userType string) ([]*models.Memo, error)
}

type memoServiceImpl struct {
	counterRepository repositories.CounterRepository
	memoRepository    repositories.MemoRepository
}

func NewMemoService(
	counterRepository repositories.CounterRepository,
	memoRepository repositories.MemoRepository,
) MemoService {
	return &memoServiceImpl{
		counterRepository: counterRepository,
		memoRepository:    memoRepository,
	}
}

// CreateMemo implements MemoService.
func (m *memoServiceImpl) CreateMemo(ctx context.Context, req *requests.MemoRequest, userRole models.Role) (*models.Memo, error) {
	// check counter exists
	counter, err := m.counterRepository.GetByName(ctx, string(userRole))
	if err != nil {
		return nil, err

	}

	if counter == nil {
		counter, err = m.counterRepository.Create(ctx, string(userRole))
		if err != nil {
			return nil, err
		}
	}

	memo := &models.Memo{
		Content:   req.Content,
		UserType:  userRole,
		CreatedAt: time.Now(),
	}

	if userRole == models.AdminRole {
		memo.Title = fmt.Sprintf("ADMIN-%d", counter.Value+1)
	} else if userRole == models.UserRole {
		memo.Title = fmt.Sprintf("MEMO-%d", counter.Value+1)
	}

	createdMemo, err := m.memoRepository.Create(ctx, memo)
	if err != nil {
		return nil, err
	}

	// Increment the counter
	err = m.counterRepository.Increment(ctx, string(userRole))
	if err != nil {
		return nil, err
	}

	return createdMemo, nil

}

// GetAllMemos implements MemoService.
func (m *memoServiceImpl) GetAllMemos(ctx context.Context) ([]*models.Memo, error) {
	memos, err := m.memoRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return memos, nil
}

// GetMemoByUserType implements MemoService.
func (m *memoServiceImpl) GetMemoByUserType(ctx context.Context, userType string) ([]*models.Memo, error) {
	// Get memos by user type
	memos, err := m.memoRepository.GetByUserType(ctx, userType)
	if err != nil {
		return nil, err
	}

	return memos, nil
}
