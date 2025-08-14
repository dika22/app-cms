package structs


type RequestCreateArticle struct {
	ArticleCategoryID int64    `json:"article_category_id" validate:"required"`
	AuthorID          int64    `json:"author_id" validate:"required"`
	Title             string   `json:"title" validate:"required"`
	Content           string   `json:"content" validate:"required"`
	TagNames          []string `json:"tags" validate:"required"`
}