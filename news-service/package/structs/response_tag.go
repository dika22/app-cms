package structs

type ResponseTag struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	UsageCount    int64  `json:"usage_count"`
	TrendingScore float64 `json:"trending_score"`
}


type ResponseTagName struct {
	Name string `json:"name"`
}