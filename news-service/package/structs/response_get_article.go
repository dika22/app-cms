package structs

type ResponseGetArticle struct {
	Page     int `json:"page"`
	Limit    int `json:"limit"`
	Total    int `json:"total"`
	ResponseArticleWithVersion []*ResponseArticleWithVersion `json:"articles"`
}