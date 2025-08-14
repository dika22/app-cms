package structs

type DetailTag struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	UsageCount    int    `json:"usage_count"`
	TrendingScore int    `json:"trending_score"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}