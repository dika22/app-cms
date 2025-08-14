package structs

import (
	"time"
)

type Article struct {
	ID       		          int64        `json:"id"`
	ArticleCategoryID         int64        `json:"article_category_id"`
	AuthorID 	              int64        `json:"author_id"`
	CurrentPublishedVersionID int64
	CurrentVersion   		  []*ArticleVersion `json:"current_version"`
	LatestVersionID           int64        `json:"latest_version_id"`
	CreatedAt 	              time.Time    `json:"created_at"`
	UpdatedAt 	              time.Time    `json:"updated_at"`
	CreatedBy                 int64        `json:"created_by"`
	UpdatedBy                 int64        `json:"updated_by"`
}


func (p RequestCreateArticle) NewArticle() Article {
	return Article{
		AuthorID:   p.AuthorID,
		ArticleCategoryID: p.ArticleCategoryID,
		LatestVersionID: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		CreatedBy:  p.AuthorID,
		UpdatedBy:  p.AuthorID,
	}
}

func (r RequestCreateArticle) NewArticleVersion(articleID int64, tags []*Tag) ArticleVersion {
	return ArticleVersion{
		ArticleID: articleID,
		VersionNumber: 1,
		Title: r.Title,
		ContentHTML: r.Content,
		Status: 1,
		ArticleTagRelationshipScore: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		CreatedBy:  r.AuthorID,
		UpdatedBy:  r.AuthorID,
		Tags: tags,
	}
	
}