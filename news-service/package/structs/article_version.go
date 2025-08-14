package structs

import "time"

type ArticleVersion struct {
	ID                          int64          `gorm:"primaryKey"`
	ArticleID                   int64          `gorm:"index;not null"`
	VersionNumber               int64          `gorm:"not null"`
	Title                       string         `gorm:"type:text;not null"`
	ContentHTML                 string         `gorm:"type:text;not null"`
	Status 						int64  		   `gorm:"index:idx_article_versions_status_id,priority:1"` // index part 1, ex val: 1 draft, 2 published 3 archived 4 deleted
	ArticleTagRelationshipScore float64        `gorm:"default:0"`
	CreatedAt                   time.Time
	PublishedAt                 time.Time
	UpdatedAt                   time.Time
	CreatedBy                   int64		   `gorm:"default:0"`
	UpdatedBy                   int64		   `gorm:"default:0"`
	Tags                        []*Tag         `gorm:"many2many:article_version_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tags"`
}


func (r RequestCreateArticleVersion) NewArticleVersionWithArticleID(nextVersion int64, tags []*Tag) ArticleVersion {
	return ArticleVersion{
		ArticleID: r.ArticleID,
		VersionNumber: nextVersion,
		Title: r.Title,
		ContentHTML: r.ContentHTML,
		Status: 1,
		ArticleTagRelationshipScore: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		CreatedBy:  r.AuthorID,
		UpdatedBy:  r.AuthorID,
		Tags: tags,
	}
	
}