package repository

import (
	"context"
	"database/sql"
	"news-service/package/structs"
)

func (r ArticleRepository) GetByIDWithVersion(ctx context.Context, id int64) (*structs.ResponseArticleWithVersion, error) {
	var result structs.ResponseArticleWithVersion
    err := r.db.Raw(`
        SELECT a.id AS article_id,
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
                   ORDER BY av.version_number DESC
                 ) FILTER (WHERE av.id IS NOT NULL), '[]'
               ) AS versions
        FROM articles a
        LEFT JOIN article_versions av ON av.article_id = a.id
        WHERE a.id = @article_id
        GROUP BY a.id, a.author_id, a.article_category_id, a.created_at, a.created_by, a.updated_by
        ORDER BY a.created_at DESC
    `, sql.Named("article_id", id)).Scan(&result).Error

    if err != nil {
        return nil, err
    }

    return &result, nil
}