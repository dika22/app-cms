package usecase

import (
	"context"

	"news-service/package/structs"
)

func (u *ArticleUsecase) GetAll(ctx context.Context, req structs.RequestSearchArticle) (structs.ResponseGetArticle, error) {
	if req.Page < 1 {
        req.Page = 1
    }
    if req.Limit <= 0 {
        req.Limit = 10 // default limit
    }
	res, err := u.repo.GetArticlesWithVersions(ctx, req)
	if err != nil {
		return structs.ResponseGetArticle{}, err
	}

	return structs.ResponseGetArticle{
		Total: len(res),
		Page:  req.Page,
		Limit: req.Limit,
		ResponseArticleWithVersion: res,
	}, nil
}