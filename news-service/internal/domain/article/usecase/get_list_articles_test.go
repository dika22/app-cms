package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"news-service/internal/domain/article/repository/mocks"
	authorMocks "news-service/internal/domain/author/repository/mocks"
	"news-service/package/structs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	cacheMock "news-service/internal/domain/article/repository/cache/mocks"
)

// -------------------- TEST ------------------------
func TestGetAll_Success(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.IRepository)
	mockAuthorRepo := new(authorMocks.IRepository)
	mockCache := new(cacheMock.CacheRepository)

	req := structs.RequestSearchArticle{
		Keyword: "golang",
	}

	expectedResp := &structs.ResponseGetArticle{
		Total: 1,
		ResponseArticleWithVersion: []*structs.ResponseArticleWithVersion{
			{
				ArticleID:   1,
				AuthorID:    1,
				CategoryID:  1,
				CreatedAt:   time.Now(),
				CreatedBy:   1,
				UpdatedBy:   1,
				Versions:    []byte("[]"),
				Rank:        0,
			},
		},
	}

	// Step 1: Cache miss
	mockCache.On("Get", ctx, req, mock.Anything).Return(errors.New("cache miss"))
	// Mocking Cache
	mockCache.On("Get", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("cache miss"))
	mockCache.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	u := &ArticleUsecase{
		db:         nil,
		avRepo:     nil,
		repo:       mockRepo,
		authorRepo: mockAuthorRepo,
		cache:      mockCache,
	}

	result, err := u.GetAll(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.Total, result.Total)
	assert.Equal(t, expectedResp.ResponseArticleWithVersion[0].ArticleID, result.ResponseArticleWithVersion[0].ArticleID)

	mockCache.AssertExpectations(t)
}
