package structs


type RequestUpdatePublishArticle struct {
	ID                		  int64  `json:"id"`
	LatestVersion    		  int64  `json:"latest_version"`
	AuthorID          		  int64  `json:"author_id" validate:"required"`
	Status                    int64  `json:"status" validate:"required"`
}