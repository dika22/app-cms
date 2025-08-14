package delivery

import (
	"net/http"

	"news-service/internal/domain/tags/usecase"
	"news-service/package/response"
	"news-service/package/structs"
	"news-service/package/validator"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type TagHTTPHandler struct {
	usecase usecase.ITags
	v *validator.Validator
}

// CreateTag godoc
// @Summary      Create new tag
// @Description  Create tag with name
// @Tags         tags
// @Accept       json
// @Produce      json
// @Param        request body structs.RequestCreateTag true "Tag to create"
// @Success      201  {object}  structs.Response
// @Router /api/v1/tags [post]
func (h *TagHTTPHandler) CreateTag(c echo.Context) error {
	var reqTag structs.RequestCreateTag
	if err := c.Bind(&reqTag); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}

	if err := h.v.UnmarshallJSONValidate(c, reqTag); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}

	if err := h.usecase.Create(c.Request().Context(), reqTag); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, reqTag, "tag created successfully")
}


// GetAllTags godoc
// @Summary      Get all tags
// @Description  Get all tags
// @Tags         tags
// @Accept       json
// @Produce      json
// @Success      200  {object}  structs.Response
// @Router /api/v1/tags [get]
func (h *TagHTTPHandler) GetAllTags(c echo.Context) error {
	tags, err := h.usecase.GetAll(c.Request().Context())
	if err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, tags, "success get all tags")
}

// GetTagByID godoc
// @Summary      Get tag by id
// @Description  Get tag with id
// @Tags         tags
// @Accept       json
// @Produce      json
// @Param        id path int true "Tag ID"
// @Success      200      {object}  structs.Response
// @Router       /api/v1/tags/:id [get]
func (h *TagHTTPHandler) GetTagByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.JSONResponse(c, http.StatusBadRequest, false, "id required", nil)
	}
	tag, err := h.usecase.GetByID(c.Request().Context(), cast.ToInt64(id))
	if err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, tag, "success get tag by id")
}

// DeleteTagByID godoc
// @Summary      Delete tag by id
// @Description  Delete tag with id
// @Tags         tags
// @Accept       json
// @Produce      json
// @Param        id path int true "Tag ID"
// @Success      200      {object}  structs.Response
// @Router       /api/v1/tags/:id [delete]
func (h TagHTTPHandler) DeleteTagByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.JSONResponse(c, http.StatusBadRequest, false, "id required", nil)
	}
	if err := h.usecase.DeleteByID(c.Request().Context(), cast.ToInt64(id)); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, nil, "success delete tag by id")
}

// UpdateTag godoc
// @Summary      Update tag by id
// @Description  Update tag with id
// @Tags         tags
// @Accept       json
// @Produce      json
// @Param        id path int true "Tag ID"
// @Param        request body structs.RequestUpdateTag true "Tag to update"
// @Success      200      {object}  structs.Response
// @Router       /api/v1/tags/:id [put]
func (h TagHTTPHandler) UpdateTag(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.JSONResponse(c, http.StatusBadRequest, false, "id required", nil)
	}
	var reqTag structs.RequestUpdateTag
	if err := c.Bind(&reqTag); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	if err := h.v.UnmarshallJSONValidate(c, reqTag); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	reqTag.ID = cast.ToInt64(id)
	if err := h.usecase.Update(c.Request().Context(), reqTag); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, reqTag, "success update tag by id")
	
}

func NewTagHTTP(r *echo.Group, uc usecase.ITags, v *validator.Validator)  {
	u := TagHTTPHandler{usecase: uc, v: v}
	r.GET("/tags", u.GetAllTags).Name = "tag.get-all"
	r.POST("/tags", u.CreateTag).Name = "tag.create"
	r.PUT("/tags/:id", u.UpdateTag).Name = "tag.update"
	r.GET("/tags/:id", u.GetTagByID).Name = "tag.get-by-id"
}