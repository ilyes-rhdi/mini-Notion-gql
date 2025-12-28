package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ilyes-rhdi/buildit-Gql/internal/services"
	"github.com/ilyes-rhdi/buildit-Gql/pkg/types"
	"github.com/labstack/echo/v4"
)

type profileHandler struct {
	srv *services.ProfileService
}

func NewProfileHandler() *profileHandler {
	return &profileHandler{
		srv: services.NewProfileService(),
	}
}

// @Summary	Get Profile endpoint
// @Tags		profiles
// @Accept		json
// @Produce	json
// @Success	200
// @Router		/profiles/get/:id [get]
func (h *profileHandler) Get(c echo.Context) error {
	id := c.Param("id")
	user, err := h.srv.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// @Summary	Search Profile endpoint
// @Tags		profiles
// @Accept		json
// @Produce	json
// @Param		email		query		string	false	"example@gmail.com"
// @Param		full_name	query		string	false	"aymen charfaoui"
// @Success	200			{object}	string
// @Router		/profiles/search [get]
func (h *profileHandler) Search(c echo.Context) error {
	email := c.QueryParam("email")
	name := c.QueryParam("name")

	if email == "" && name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "at least provide one query param")
	}

	if email != "" {
		user, err := h.srv.GetUserByEmail(email)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}

	users, err := h.srv.SearchByName(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// @Summary	Change Profile Image endpoint
// @Tags		profiles
// @Accept		form/multipart
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		image	formData	file	true	"file.png"
// @Success	200
// @Router		/profiles/profile/pfp [patch]
func (h *profileHandler) ChangePfp(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*types.Claims)

	// Ensure directory exists
	dir := "public/uploads/profiles"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	filename := filepath.Base(file.Filename)
	path := fmt.Sprintf("%s/%s", dir, filename)

	dst, err := os.Create(path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	defer dst.Close()

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	defer src.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedPath, err := h.srv.UpdateUserImage(claims.ID, path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, types.Response{
		"image": updatedPath,
	})
}

// @Summary	Change Profile Bg Image endpoint
// @Tags		profiles
// @Accept		form/multipart
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		image	formData	file	true	"file.png"
// @Success	200
// @Router		/profiles/profile/bg [patch]
func (h *profileHandler) ChangeBg(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*types.Claims)

	// Ensure directory exists
	dir := "public/uploads/bgs"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	filename := filepath.Base(file.Filename)
	path := fmt.Sprintf("%s/%s", dir, filename)

	dst, err := os.Create(path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	defer dst.Close()

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	defer src.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedPath, err := h.srv.UpdateUserBg(claims.ID, path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, types.Response{
		"bgImg": updatedPath,
	})
}

// @Summary	Update Profile endpoint
// @Tags		profiles
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		body	body	types.ProfileUpdate	false	"jhon doe"
// @Success	200
// @Router		/profiles/profile/update [patch]
func (h *profileHandler) Update(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*types.Claims)

	u, err := h.srv.GetUser(claims.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Avec GORM, ces champs sont normalement des strings directes (pas des optionnels Prisma).
	payload := types.ProfileUpdate{
		Email:  u.Email,
		Name:   u.Name,
		Bio:    u.Bio,
	}

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updated, err := h.srv.UpdateUser(claims.ID, payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, updated)
}

// @Summary	Get Current Profile endpoint
// @Tags		profiles
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Success	200
// @Router		/profiles/profile [get]
func (h *profileHandler) CurrentUser(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*types.Claims)

	u, err := h.srv.GetUser(claims.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, u)
}

// @Summary	Delete Profile endpoint
// @Tags		profiles
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Success	200
// @Router		/profiles/profile/delete [delete]
func (h *profileHandler) Delete(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*types.Claims)

	deletedID, err := h.srv.DeleteUser(claims.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, types.Response{
		"message": fmt.Sprintf("user deleted id : %s", deletedID),
	})
}
