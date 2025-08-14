package delivery

import (
	"net/http"
	"news-service/cmd/middleware"
	"news-service/internal/domain/article-version/usecase"
	"news-service/package/response"
	"news-service/package/structs"
	"news-service/package/validator"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type ArticleVersionHTTP struct {
	uc usecase.IArticleVersionUsecase
	v *validator.Validator
}

// GetAll gets all article versions
//
// @Summary Get all article versions
// @Description get all article versions
// @Tags article-versions
// @Produce json
// @Success      200      {object}  structs.Response
// @Router /api/v1/articles/versions [get]
func (h ArticleVersionHTTP) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	articleVersions, err := h.uc.ListArticleVersion(ctx)
	if err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, articleVersions, "success get all article versions")
	
}


// CreateArticleVersion godoc
// @Summary      Create new article version
// @Description  Create new article version with title and body
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param        id path int true "Article ID"
// @Param        request body structs.RequestCreateArticleVersion true "Article version to create"
// @Success      201      {object}  structs.Response
// @Router       /api/v1/articles/{id}/versions [put]
func (h ArticleVersionHTTP) UpdateArticleVersion(ctx echo.Context) error {
	id := ctx.Param("article_id")
	if id == "" {
		return response.JSONResponse(ctx, http.StatusBadRequest, false, "id required", nil)
	}
	req := structs.RequestCreateArticleVersion{}
	if err := h.v.UnmarshallJSONValidate(ctx, &req); err != nil {
		return response.JSONResponse(ctx, http.StatusBadRequest, false, err.Error(), nil)
	}
	req.ArticleID = cast.ToInt64(id)
	if err := h.uc.UpdateArticleVersion(ctx.Request().Context(), req); err != nil {
		return response.JSONResponse(ctx, http.StatusBadRequest, false, err.Error(), nil)
	}
	return nil
}

// ListArticleVersionsByArticleID retrieves all versions of an article by its ID.
//
// @Summary      List article versions by article ID
// @Description  Get all versions of a specific article using its ID
// @Tags         article-versions
// @Produce      json
// @Param        article_id path int true "Article ID"
// @Success      200 {object} structs.Response
// @Router       /api/v1/articles/{article_id}/versions [get]
func (h ArticleVersionHTTP) ListArticleVersionsByArticleID(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("article_id")
	if id == "" {
		return response.JSONResponse(c, http.StatusBadRequest, false, "id required", nil)
	}
	res, err := h.uc.GetArticleVersionByArticleID(ctx, cast.ToInt64(id))
	if err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, res.ArticleVersionDetail, "success get article versions by article id")
}

// GetArticleVersionByID godoc
// @Summary      Get article version by id
// @Description  Get article version with id and version number
// @Tags         article-versions
// @Accept       json
// @Produce      json
// @Param        id path int true "Article ID"
// @Param        version_number path int true "Version number"
// @Success      200      {object}  structs.Response
// @Router       /api/v1/articles/versions/{id}/{version_number} [get]
func (h ArticleVersionHTTP) GetArticleVersionByID(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return response.JSONResponse(ctx, http.StatusBadRequest, false, "id required", nil)
	}
	versionNumber := ctx.Param("version_number")
	if versionNumber == "" {
		return response.JSONResponse(ctx, http.StatusBadRequest, false, "version number required", nil)
	}
	res, err := h.uc.GetArticleVersionByIDAndVersion(ctx.Request().Context(), cast.ToInt64(id), cast.ToInt64(versionNumber))
	if err != nil {
		return response.JSONResponse(ctx, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(ctx, res, "success get article version by id and version number")
}

// GetArticleVersionByArticleID godoc
// @Summary      Get article version by article ID and version number
// @Description  Retrieves a specific version of an article using the article ID and version number
// @Tags         article-versions
// @Accept       json
// @Produce      json
// @Param        article_id path int true "Article ID"
// @Param        version_number path int true "Version number"
// @Success      200      {object}  structs.Response
// @Router       /api/v1/articles/versions/{article_id}/{version_number} [get]
func (h ArticleVersionHTTP) GetArticleVersionByArticleID(ctx echo.Context) error {
	id := ctx.Param("article_id")
	if id == "" {
		return response.JSONResponse(ctx, http.StatusBadRequest, false, "id required", nil)
	}

	versionNumber := ctx.Param("version_number")
	if versionNumber == "" {
		return response.JSONResponse(ctx, http.StatusBadRequest, false, "version number required", nil)
	}
	res, err := h.uc.GetArticleVersionByIDAndVersion(ctx.Request().Context(), cast.ToInt64(id), cast.ToInt64(versionNumber))
	if err != nil {
		return response.JSONResponse(ctx, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(ctx, res, "success get article version by id and version number")
}

func NewArticleVersionHTTP(r *echo.Group, uc usecase.IArticleVersionUsecase, v *validator.Validator)  {
	u := ArticleVersionHTTP{uc: uc, v: v}
	r.GET("/articles/versions", u.GetAll, middleware.AuthMiddleware).Name = "article-version.get-all"
	r.GET("/articles/:article_id/versions", u.ListArticleVersionsByArticleID, middleware.AuthMiddleware).Name = "article-version.list-by-article-id"
	r.PUT("/articles/:article_id/versions", u.UpdateArticleVersion, middleware.AuthEditorMiddleware).Name = "article-version.create"
	r.GET("/articles/:article_id/versions/:version_number", u.GetArticleVersionByArticleID, middleware.AuthMiddleware).Name = "article-version.get-by-article-id"
}