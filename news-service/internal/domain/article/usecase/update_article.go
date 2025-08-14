package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"news-service/internal/constant"
	"news-service/package/structs"

	"gorm.io/gorm"
)


func (u *ArticleUsecase) UpdatePublishArticle(ctx context.Context, req *structs.RequestUpdatePublishArticle) error {
	articleVersion := structs.ArticleVersion{}
	if err := u.avRepo.GetArticleVersionByArticleIDAndVersion(ctx, structs.RequestGetArticle{ID: req.ID, Version: req.LatestVersion}, &articleVersion); err != nil {
		return err
	}

	oldStatus := articleVersion.Status
	if oldStatus == req.Status {
		msg := fmt.Sprintf("article is already %v", constant.ArticleStatus[int(req.Status)])
		return errors.New(msg)
	}

	newStatus := articleVersion.Status
	if req.Status == constant.Published {
		now := time.Now()
		articleVersion.PublishedAt = now
		articleVersion.UpdatedAt = now
		articleVersion.UpdatedBy = req.AuthorID
		articleVersion.Status = constant.Published
		newStatus = constant.Published
		articleVersion.ContentHTML = req.ContentHTML
		articleVersion.Title = req.Title
	}

	if err := u.avRepo.UpdateByArticleIDAndVersion(ctx, &articleVersion); err != nil {
		return err
	}

	article := structs.Article{}
	if err := u.repo.GetByID(ctx, req.ID, &article); err != nil {
		return err
	}

	if article.ID < 1 {
		return errors.New("article not found")
	}

	if req.ID != article.ID {
		return errors.New("article id not match")
		
	}

	article.CurrentPublishedVersionID = articleVersion.ID
	article.UpdatedAt = time.Now()
	article.UpdatedBy = req.AuthorID
	_, err := u.repo.Update(ctx, &article)
	if err != nil {
		return err
	}

	if err := handleTagUsageStatusChange(u.db, &articleVersion, int64(oldStatus), int64(newStatus)); err != nil {
		return err
	}
	
	return nil
}

func handleTagUsageStatusChange(db *gorm.DB, av *structs.ArticleVersion, oldStatus, newStatus int64) error {
	// Published -> Unpublished (decrement)
	if oldStatus == constant.Published && newStatus != constant.Published  {
		if err := updateTagUsageCounts(db, av.Tags, -1); err != nil {
			db.Rollback()
			return err
		}
	}

	// Unpublished -> Published (increment)
	if oldStatus != constant.Published && newStatus == constant.Published {
		if err := updateTagUsageCounts(db, av.Tags, +1); err != nil {
			db.Rollback()
			return err
		}
	}

	return nil
}

func updateTagUsageCounts(db *gorm.DB, tags []*structs.Tag, delta int64) error {
	for _, t := range tags {
		t.UsageCount += delta
		t.TrendingScore = computeTrending(t)
		if err := db.Save(t).Error; err != nil {
			return err
		}
	}
	return nil
}

func computeTrending(t *structs.Tag) float64 {
	// naive: usage_count / days since created + 1
	days := time.Since(t.CreatedAt).Hours() / 24.0
	if days < 1 { days = 1 }
	return float64(t.UsageCount) / (1.0 + days)
}
