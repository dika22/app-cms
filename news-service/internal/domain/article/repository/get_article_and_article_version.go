package repository

import (
	"context"
	"news-service/package/structs"
	"strings"
)



func (r ArticleRepository) GetArticlesWithVersions(ctx context.Context, p structs.RequestSearchArticle) ([]*structs.ResponseArticleWithVersion, error) {
	var results []*structs.ResponseArticleWithVersion
    offset := (p.Page - 1) * p.Limit
    baseQuery := r.db.Table("articles a").
    Select(`
        a.id AS article_id,
        a.author_id,
        a.article_category_id,
        a.created_at,
        a.created_by,
        a.updated_by,
        COALESCE(
            json_agg(
                json_build_object(
                    'id', av.id,
                    'version_number', av.version_number,
                    'status', av.status,
                    'title', av.title,
                    'content_html', av.content_html,
                    'article_tag_relationship_score', av.article_tag_relationship_score
                )
            ) FILTER (WHERE av.id IS NOT NULL), '[]'
        ) AS versions,
        CASE 
            WHEN ? <> '' THEN MAX(ts_rank_cd(
                to_tsvector('simple', COALESCE(av.title, '') || ' ' || COALESCE(av.content_html, '')),
                plainto_tsquery('simple', ?)
            ))
            ELSE 0
        END AS rank
    `, p.Keyword, p.Keyword).
    Joins("LEFT JOIN article_versions av ON av.article_id = a.id").
    Group("a.id, a.author_id, a.article_category_id, a.created_at, a.created_by, a.updated_by")

    // Jika search tidak kosong → tambahkan pencarian dan urutan relevansi
    if strings.TrimSpace(p.Keyword) != "" {
        baseQuery = baseQuery.Where(
            `to_tsvector('simple', COALESCE(av.title, '') || ' ' || COALESCE(av.content_html, '')) @@ plainto_tsquery('simple', ?)`,
            p.Keyword,
        ).Order("rank DESC").Order("a.created_at DESC")
    } else {
        baseQuery = baseQuery.Order("a.created_at DESC")
    }

    // Pagination
    baseQuery = baseQuery.Offset(offset).Limit(p.Limit)
    if err := baseQuery.Scan(&results).Error; err != nil {
        return nil, err
    }

    return results, nil
}
