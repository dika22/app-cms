package usecase

import (
	"context"

	"news-service/package/structs"

	"gorm.io/gorm"
)

func (u *ArticleUsecase) Create(ctx context.Context, req *structs.RequestCreateArticle) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		article := req.NewArticle()
		arctileID, err := u.repo.Store(ctx, article)
		if err != nil {
			return err
		}

		tags := []*structs.Tag{}
		if err := u.tagRepo.FindOrCreateTagsByNames(ctx, req.TagNames, &tags); err != nil {
			return err
		}

		// Buat set untuk tag yang sudah ditemukan
		existing := make(map[string]struct{}, len(tags))
		for _, t := range tags {
			existing[t.Name] = struct{}{}
		}

		// Cari tag yang belum ada
		var notFound []string
		for _, name := range req.TagNames {
			if _, ok := existing[name]; !ok {
				notFound = append(notFound, name)
			}
		}

		// Simpan tag yang belum ada
		for _, name := range notFound {
			if err := u.tagRepo.Store(ctx, structs.Tag{Name: name}); err != nil {
				return err
			}
		}

		tagIDs := make([]int64, 0, len(tags))
		for _,t :=range tags{
			tagIDs = append(tagIDs, t.ID) 
		}

		score := u.ComputeArticleTagRelationshipScore(tagIDs)
		articleVersion := req.NewArticleVersion(arctileID, tags)
		articleVersion.ArticleTagRelationshipScore = score
		if err = u.avRepo.CreateArticleVersion(ctx, articleVersion); err != nil {
			return err
		}

		_, err = u.repo.UpdateVersioning(ctx, &structs.RequestUpdateVersioning{
			ID: article.ID,
			Version: articleVersion.VersionNumber,
		})
		if err != nil {
			return err
		}
		return nil
	})
}


// computeArticleTagRelationshipScore menghitung skor relasi antar tag dalam 1 artikel.
// Skor ini dihitung dengan cara:
// 1. Hitung occurrence setiap tag
// 2. Hitung co-occurrence setiap pasangan tag
// 3. Hitung rata-rata skor dengan rumus:
// score = count(pasangan) / (occurrence(tagA) + occurrence(tagB) - count(pasangan))
func (u *ArticleUsecase) ComputeArticleTagRelationshipScore(tagIDs []int64) float64 {
	if len(tagIDs) < 2 {
		return 0
		
	}
	// 1. Hitung semua occurrence tag sekaligus
    var occs []structs.CountResult
	if err  := u.tagRepo.CountOccuranceAllTags(context.Background(), tagIDs, &occs); err != nil {
		return 0
	}

    occMap := make(map[int64]int64, len(occs))
    for _, o := range occs {
        occMap[o.TagID] = o.Count
    }

    // 2. Hitung semua co-occurrence pasangan sekaligus
    var cos []structs.CoResult
	if err := u.tagRepo.CountOccuranceCoupleTags(context.Background(), tagIDs, &cos); err != nil {
		return 0
	}

    // 3. Hitung skor rata-rata
    sum := 0.0
    pairs := 0
    for _, c := range cos {
        occA := occMap[c.TagA]
        occB := occMap[c.TagB]
        den := occA + occB - c.Count
        var score float64
        if den > 0 {
            score = float64(c.Count) / float64(den)
        }
        sum += score
        pairs++
    }

    if pairs == 0 {
        return 0
    }
    return sum / float64(pairs)
}
