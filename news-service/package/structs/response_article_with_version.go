package structs

import (
	"encoding/json"
	"time"
)

type ResponseArticleWithVersion struct {
	ArticleID   int64           `json:"article_id"`
    AuthorID    int64           `json:"author_id"`
    CategoryID  int64           `json:"article_category_id"`
    CreatedAt   time.Time       `json:"created_at"`
    CreatedBy   int64           `json:"created_by"`
    UpdatedBy   int64           `json:"updated_by"`
    Versions    json.RawMessage `json:"versions"`
    Rank        float64         `json:"rank"`
}


type Versions struct {
	ID                          int    `json:"id"`
	VersionNumber               int    `json:"version_number"`
	Status                      int    `json:"status"`
	Title                       string `json:"title"`
	ContentHTML                 string `json:"content_html"`
	ArticleTagRelationshipScore int    `json:"article_tag_relationship_score"`
}