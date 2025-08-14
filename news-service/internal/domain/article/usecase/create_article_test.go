package usecase_test

import (
	"context"
	"testing"

	"news-service/internal/domain/article/repository/mocks"
	"news-service/internal/domain/article/usecase"
	authorMocks "news-service/internal/domain/author/repository/mocks"
	"news-service/package/structs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate_Success(t *testing.T) {
	// Mock Dependencies
	mockRepo := new(mocks.IRepository)
	mockAuthorRepo := new(authorMocks.IRepository)

	usecase := usecase.NewsUsecase(
		nil,
		mockRepo,
		mockAuthorRepo,
		nil,
		nil,
		nil,
		nil,
	)

	req := &structs.RequestCreateArticle{
		Title:    "Judul",
		Content:     "Konten",
		AuthorID: 1,
	}

	article := req.NewArticle()
	article.ID = 1

	mockRepo.On("Store", mock.Anything, mock.Anything).Return(int64(1), nil)
	mockAuthorRepo.On("GetByID", mock.Anything, int64(1)).Return(structs.Authors{
		ID:   1,
		Name: "Adhika",
	}, nil)

	err := usecase.Create(context.Background(), req)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockAuthorRepo.AssertExpectations(t)
}

