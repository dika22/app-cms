package structs

type ResponseArticleVersionByArticleID struct {
	ArticleVersionDetail []ArticleVersionDetail `json:"article_version_detail"`
}

type ArticleVersionDetail struct {
	ID                          int64          `json:"id"`
	ArticleID                   int64          `json:"article_id"`
	VersionNumber               int64          `json:"version_number"`
	Title                       string         `json:"title"`
	ContentHTML                 string         `json:"content_html"`
	Status 						int64  		   `json:"status"` // index part 1, ex val: 1 draft, 2 published 3 archived 4 deleted
	ArticleTagRelationshipScore float64        `json:"article_tag_relationship_score"`
	CreatedAt                   string         `json:"created_at"`
	CreatedBy                   int64		   `json:"created_by"`
	UpdatedBy                   int64		   `json:"updated_by"`
	Tags                        []string		`json:"tags"`
}

func (av ArticleVersion) NewArticleVersionDetail() ArticleVersionDetail {
	return ArticleVersionDetail{
		ID: av.ID,
		ArticleID: av.ArticleID,
		VersionNumber: av.VersionNumber,
		Status: av.Status,
		CreatedAt: av.CreatedAt.String(),
		CreatedBy: av.CreatedBy,
		UpdatedBy: av.UpdatedBy,
		Title: av.Title,
		ContentHTML: av.ContentHTML,
		ArticleTagRelationshipScore: av.ArticleTagRelationshipScore,
		Tags: av.MapTags(),
	}	
}

func (av ArticleVersion) MapTags() []string {
	tags := []string{}
	for _, tag := range av.Tags {
		tags = append(tags, tag.Name)
	}
	return tags
}