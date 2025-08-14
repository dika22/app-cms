package usecase

import (
	"context"
	"news-service/package/structs"
)

func (u *ArticleVersionUsecase) GetArticleVersionByArticleIDAndVersion(ctx context.Context, id int64, versionNumber int64) (*structs.ArticleVersionDetail, error){
	dest := structs.ArticleVersion{} 
	err := u.repo.GetArticleVersionByArticleIDAndVersion(ctx, structs.RequestGetArticle{ID: id, Version: versionNumber}, &dest)
	if err != nil {
		return nil, err
	}
	avDetail := dest.NewArticleVersionDetail()
	return &avDetail, nil
}