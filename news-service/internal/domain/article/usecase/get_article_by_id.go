package usecase

import (
	"context"

	"news-service/package/structs"
)


func (u *ArticleUsecase) GetByID(ctx context.Context, id int64) (structs.ResponseArticleDetail, error) {
	// res, err, _ := u.group.Do(cast.ToString(id), func() (interface{}, error) {
	// 	if err := u.repo.GetByID(ctx, id, dest); err != nil {
	// 		return structs.Article{}, err
	// 	}
	// 	return dest, nil
    // })
	// if err != nil {
	// 	return structs.ResponseArticle{}, err
	// }
	
	article := structs.Article{}
	if err := u.repo.GetByID(ctx, id, &article); err != nil {
		return structs.ResponseArticleDetail{}, err
	}

	articleVersion := structs.ArticleVersion{}
	if err := u.avRepo.GetArticleVersionByArticleID(ctx, article.ID, &articleVersion); err != nil {
		return structs.ResponseArticleDetail{}, err
	}

	author, err := u.authorRepo.GetByID(ctx, int64(article.AuthorID))
	if err != nil {
		return structs.ResponseArticleDetail{}, err
	}

	resp := structs.ResponseArticleDetail{
		ArticleID: article.ID,
		AuthorID: article.AuthorID,
		CategoryID: article.ArticleCategoryID,
		CreatedAt: article.CreatedAt,
		CreatedBy: article.CreatedBy,
		UpdatedBy: article.UpdatedBy,
		ArticleWithVersion: []structs.ArticleVersionDetail{articleVersion.NewArticleVersionDetail()},
		Author: structs.Author{ID: author.ID, Name: author.Name},
	}

	return  resp, nil
}