package usecase

import (
	"context"
	"news-service/package/structs"
)

func (u *ArticleVersionUsecase) ListArticleVersion(ctx context.Context) ([]structs.ArticleVersion, error){
	dest := []structs.ArticleVersion{}
	if err := u.repo.GetAll(ctx, &dest); err != nil {
		return nil, err
	}
	return dest, nil
	
}