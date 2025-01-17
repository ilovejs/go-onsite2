package handler

import (
	"github.com/labstack/echo"
	"net/http"
	. "onsite/dto/users"
	"onsite/models"
	. "onsite/utils"
	"strconv"
)

func (h *Handler) SignUp(c echo.Context) error {
	var u models.User
	req := &UserRegisterRequest{}
	if err := req.Bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	// create user
	if err := h.userStore.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	// spew.Dump(u)
	return c.JSON(http.StatusCreated, NewUserResponse(&u)) // new token
}

func (h *Handler) Login(c echo.Context) error {
	req := &UserLoginRequest{}

	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	u, err := h.userStore.GetByEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, AccessForbidden())
	}
	if !CheckPassword(u, req.Password) {
		return c.JSON(http.StatusForbidden, H{"message": "wrong password"})
	}
	return c.JSON(http.StatusOK, NewUserLoginResponse(u))
}

func (h *Handler) Logout(c echo.Context) error {
	c.Set("user", "")
	return c.JSON(http.StatusOK, NewUserLogoutResponse())
}

func (h *Handler) CurrentUser(c echo.Context) error {
	u, err := h.userStore.GetByID(getUserId(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}
	return c.JSON(http.StatusOK, NewUserResponse(u))
}

func (h *Handler) UpdateUser(c echo.Context) error {
	// find user from JWT context
	targetUser, err := h.userStore.GetByID(getUserId(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if targetUser == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}

	req := NewUserUpdateRequest()
	req.LoadFrom(targetUser)

	// 1. take context to validate against request format,
	// 2. attach validated data to u (existing record object)
	if err := req.Bind(c, targetUser); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	// update
	if err := h.userStore.Update(targetUser); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	return c.JSON(http.StatusOK, NewUserResponse(targetUser)) // new token
}

func (h *Handler) ListUser(c echo.Context) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	tradeItems, count, err := h.userStore.List(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusOK, NewListUserResponse(tradeItems, count))
}

func (h *Handler) GetProfile(c echo.Context) error {
	username := c.Param("username")
	u, err := h.userStore.GetByUserName(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}
	return c.JSON(http.StatusOK, NewProfileResponse(u))
}

/* utils of handler */

func getUserId(c echo.Context) int {
	id, ok := c.Get("user").(int)
	if !ok {
		return 0
	}
	return id
}
