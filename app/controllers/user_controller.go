package controllers

import (
	"echo-template/app/models"
	"echo-template/app/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// GetUsers 获取用户列表
// @Summary      获取用户列表
// @Description  获取所有用户信息
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.UserListResponse  "成功返回用户列表"
// @Failure      500  {object}  models.ErrorResponse  "服务器错误"
// @Router       /v1/users [get]
func (uc *UserController) GetUsers(c echo.Context) error {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.UserListResponse{
		Code: http.StatusOK,
		Data: users,
		Msg:  "获取用户列表成功",
	})
}

// GetUser 获取单个用户
// @Summary      获取单个用户
// @Description  根据ID获取用户详细信息
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "用户ID"  example(1)
// @Success      200  {object}  models.UserResponse  "成功返回用户信息"
// @Failure      400  {object}  models.ErrorResponse  "请求参数错误"
// @Failure      404  {object}  models.ErrorResponse  "用户不存在"
// @Router       /v1/users/{id} [get]
func (uc *UserController) GetUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code: http.StatusBadRequest,
			Msg:  "Invalid user ID",
		})
	}

	user, err := uc.userService.GetUserByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code: http.StatusNotFound,
			Msg:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.UserResponse{
		Code: http.StatusOK,
		Data: *user,
		Msg:  "获取用户信息成功",
	})
}

// CreateUser 创建用户
// @Summary      创建用户
// @Description  创建新用户
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "用户信息"
// @Success      201   {object}  models.UserResponse  "成功创建用户"
// @Failure      400   {object}  models.ErrorResponse  "请求参数错误"
// @Failure      500   {object}  models.ErrorResponse  "服务器错误"
// @Router       /v1/users [post]
func (uc *UserController) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code: http.StatusBadRequest,
			Msg:  "Invalid request body",
		})
	}

	if err := uc.userService.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.UserResponse{
		Code: http.StatusCreated,
		Data: user,
		Msg:  "创建用户成功",
	})
}

// UpdateUser 更新用户
// @Summary      更新用户
// @Description  更新用户信息
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int          true  "用户ID"  example(1)
// @Param        user  body      models.User  true  "用户信息"
// @Success      200   {object}  models.UserResponse  "成功更新用户"
// @Failure      400   {object}  models.ErrorResponse  "请求参数错误"
// @Failure      500   {object}  models.ErrorResponse  "服务器错误"
// @Router       /v1/users/{id} [put]
func (uc *UserController) UpdateUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code: http.StatusBadRequest,
			Msg:  "Invalid user ID",
		})
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code: http.StatusBadRequest,
			Msg:  "Invalid request body",
		})
	}

	user.ID = uint(id)
	if err := uc.userService.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.UserResponse{
		Code: http.StatusOK,
		Data: user,
		Msg:  "更新用户成功",
	})
}

// DeleteUser 删除用户
// @Summary      删除用户
// @Description  根据ID删除用户
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "用户ID"  example(1)
// @Success      200  {object}  models.SuccessResponse  "成功删除用户"
// @Failure      400  {object}  models.ErrorResponse  "请求参数错误"
// @Failure      500  {object}  models.ErrorResponse  "服务器错误"
// @Router       /v1/users/{id} [delete]
func (uc *UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code: http.StatusBadRequest,
			Msg:  "Invalid user ID",
		})
	}

	if err := uc.userService.DeleteUser(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Code: http.StatusOK,
		Msg:  "删除用户成功",
	})
}
