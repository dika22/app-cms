package usecase

import (
	"context"
	"errors"
	"fmt"
	"news-service/package/structs"

	"gorm.io/gorm"
)

func (u *ArticleUsecase) DeleteArticleByID(ctx context.Context, articleID int64) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		article := structs.Article{}
		if err := u.repo.GetByID(ctx, articleID, &article); err != nil {
			fmt.Println("debug", err)
			return err
		}

		if  article.ID == 0 {
			return errors.New("article not found")
		}

		// Delete article
		if err := u.repo.Delete(ctx, articleID); err != nil { 
			return err 
		}

		// get article version
		articleVersion := structs.ArticleVersion{}
		if err := u.avRepo.GetArticleVersionByArticleID(ctx, articleID, &articleVersion); err != nil {
			return err
		}

		if articleVersion.ID == 0 {
			return errors.New("article version not found")
		}

		// Delete article version
		if err := u.avRepo.DeleteByID(ctx, articleVersion.ID); err != nil {
			return err
		}
		return nil
	})
}