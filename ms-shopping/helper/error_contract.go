package helper

import (
	"net/http"
	"phase3-gc1-shopping/model/web"
)

func ErrBadRequest(detail any) web.ErrorResponse {
	return web.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
		Detail:  detail,
	}
}

func ErrInternalServer(detail any) web.ErrorResponse {
	return web.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Detail:  detail,
	}
}
