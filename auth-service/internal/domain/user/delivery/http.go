package delivery

import (
	"auth-service/internal/domain/user/usecase"
	"auth-service/package/response"
	"auth-service/package/structs"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

type UserHTTP struct {
	uc usecase.IUser
}

// SignUp godoc
// @Summary      Create new user
// @Description  Create user with name, email, and password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body structs.RequestSignUp true "User to create"
// @Success      201  {object}  structs.Response
// @Failure      400  {object}  structs.Response
// @Router /api/v1/users [post]
func (h UserHTTP) SignUp(c echo.Context) error {
	ctx := c.Request().Context()
	req := structs.RequestSignUp{}
	if err := c.Bind(&req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	if err := h.uc.SignUp(ctx, req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, nil, "success create user")
}

// Login godoc
// @Summary      Login user
// @Description  Login user with email and password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body structs.RequestLogin true "User to login"
// @Success      200  {object}  structs.ResponseLogin
// @Failure      401  {object}  structs.Response
// @Router /api/v1/login [post]
func (h UserHTTP) Login(c echo.Context) error {
	ctx := c.Request().Context()
	req := structs.RequestLogin{}
	if err := c.Bind(&req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	resp, err := h.uc.Login(ctx, req)
	if err != nil {
		log.Error().Err(err).Str("endpoint", "/login").Msg("not authorized")
		return response.JSONResponse(c, http.StatusUnauthorized, false, err.Error(), nil)
	}
	return response.JSONResponse(c, http.StatusOK, true, "success login", resp)
}

// UpdateUser godoc
// @Summary      Update user information
// @Description  Update user details such as name, email, or password by user ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Param        request body structs.RequestUpdateUser true "User details to update"
// @Success      200  {object}  structs.Response
// @Failure      400  {object}  structs.Response
// @Router       /api/v1/users/{id} [put]

func (h UserHTTP) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	req := structs.RequestUpdateUser{}
	if err := c.Bind(&req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	req.ID = cast.ToInt64(id)
	if err := h.uc.UpdateUser(ctx, req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, nil, "success update user")
}

// GetAllTags godoc
// @Summary      Get all users
// @Description  Get all users
// @Tags         tags
// @Accept       json
// @Produce      json
// @Success      200  {object}  structs.Response
// @Router /api/v1/users/{id} [get]
func (h UserHTTP) GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	fmt.Println("id", id)
	resp, err := h.uc.GetByID(ctx, cast.ToInt64(id))
	if err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, resp, "success get user")
}

func NewUserHTTP(r *echo.Group, uc usecase.IUser) {
	u := UserHTTP{uc: uc}
	r.POST("/signup", u.SignUp).Name = "users.signup"
	r.POST("/login", u.Login).Name = "users.login"
	r.PUT("/update/:id", u.UpdateUser).Name = "update.user"
	r.GET("/:id", u.GetByID).Name = "get.by.id"
}
