package structs

type RequestCreateArticleVersion struct {
	ArticleCategoryID int64    `json:"article_category_id" validate:"required"`
	ArticleID   int64 		   `json:"article_id"`
	Title       string 		   `json:"title"`
	ContentHTML string 		   `json:"content_html"`
	AuthorID    int64          `json:"author_id"`
	TagNames    []string       `json:"tags" validate:"required"`
}