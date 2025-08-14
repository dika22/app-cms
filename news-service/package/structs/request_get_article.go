package structs

type RequestGetArticle struct {
	ID        int64 `json:"id"`
	ArticleID int64 `json:"article_id" validate:"required"`
	Version   int64 `json:"version" validate:"required"`
}