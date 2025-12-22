package models

// Response 通用响应结构
// @Description 通用API响应
type Response struct {
	Code int         `json:"code" example:"200"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg" example:"success"`
}

// SuccessResponse 成功响应（用于删除等操作）
// @Description 成功响应
type SuccessResponse struct {
	Code int    `json:"code" example:"200"`
	Msg  string `json:"msg" example:"操作成功"`
}

// ErrorResponse 错误响应
// @Description 错误响应
type ErrorResponse struct {
	Code int    `json:"code" example:"400"`
	Msg  string `json:"msg" example:"操作失败"`
}

// UserListResponse 用户列表响应
// @Description 用户列表响应
type UserListResponse struct {
	Code int    `json:"code" example:"200"`
	Data []User `json:"data"`
	Msg  string `json:"msg" example:"获取用户列表成功"`
}

// UserResponse 单个用户响应
// @Description 用户响应
type UserResponse struct {
	Code int    `json:"code" example:"200"`
	Data User   `json:"data"`
	Msg  string `json:"msg" example:"获取用户信息成功"`
}
