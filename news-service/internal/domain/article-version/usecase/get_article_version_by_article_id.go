package usecase

import (
	"context"
	"news-service/package/structs"
)

func (u *ArticleVersionUsecase) GetArticleVersionByArticleID(ctx context.Context, id int64) (*structs.ResponseArticleVersionByArticleID, error) {
	articlesVersions := []structs.ArticleVersion{}
	if err := u.repo.GetArticleVersionByArticleID(ctx, id, &articlesVersions); err != nil {
		return nil, err
	}

	mapArticleVersion := []structs.ArticleVersionDetail{}
	for _, av := range articlesVersions {
		avDetail := structs.ArticleVersionDetail{
			ID: av.ID,
			ArticleID: av.ArticleID,
			VersionNumber: av.VersionNumber,
			Status: av.Status,
			CreatedAt: av.CreatedAt.String(),
			CreatedBy: av.CreatedBy,
			UpdatedBy: av.UpdatedBy,
			Title: av.Title,
			ContentHTML: av.ContentHTML,
			ArticleTagRelationshipScore: av.ArticleTagRelationshipScore,
			Tags: av.MapTags(),
		}
		mapArticleVersion = append(mapArticleVersion, avDetail)
	}

	return &structs.ResponseArticleVersionByArticleID{
		ArticleVersionDetail: mapArticleVersion,
	}, nil
}