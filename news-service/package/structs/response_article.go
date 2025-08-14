package structs

import "time"


type ResponseArticleDetail struct {
	ArticleID   int64           `json:"article_id"`
    AuthorID    int64           `json:"author_id"`
    CategoryID  int64           `json:"article_category_id"`
    CreatedAt   time.Time       `json:"created_at"`
    CreatedBy   int64           `json:"created_by"`
    UpdatedBy   int64           `json:"updated_by"`
	ArticleWithVersion	[]ArticleVersionDetail  `json:"article"`
	Author   		    Author                `json:"author"`
}




type Author struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
}

