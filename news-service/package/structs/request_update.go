package structs


type RequestUpdatePublishArticle struct {
	ID                		  int64  `json:"id"`
	ArticleCategoryID 		  int64  `json:"article_category_id" validate:"required"`
	LatestVersion    		  int64  `json:"latest_version"`
	CurrentPublishedVersionID int64  `json:"current_published_version_id"`
	AuthorID          		  int64  `json:"author_id" validate:"required"`
	Title                     string `json:"title" validate:"required"`
	ContentHTML               string `json:"content" validate:"required"`
	Status                    int64  `json:"status" validate:"required"`
}