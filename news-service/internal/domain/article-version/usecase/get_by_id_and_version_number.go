package usecase

import (
	"context"
	"news-service/package/structs"
)

func (u *ArticleVersionUsecase) GetArticleVersionByIDAndVersion(ctx context.Context, id int64, versionNumber int64) (*structs.ArticleVersionDetail, error){
	dest := structs.ArticleVersion{} 
	if err := u.repo.GetArticleVersionByIDAndVersion(ctx, structs.RequestGetArticle{ID: id, Version: versionNumber}, &dest); err != nil {
		return nil, err
	}
	avDetail := dest.NewArticleVersionDetail()
	return &avDetail, nil
}