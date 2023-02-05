package controller

import (
	"Minimalist_TikTok/config"
	"Minimalist_TikTok/serializer"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func ErrorResponse(err error) serializer.ErrorResponse {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := config.T(fmt.Sprintf("Field.%s", e.Field))
			tag := config.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.ErrorResponse{
				StatusCode: 40001,
				StatusMsg:  fmt.Sprintf("%s%s", field, tag),
				Error:      fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ErrorResponse{
			StatusCode: 40001,
			StatusMsg:  "JSON类型不匹配",
			Error:      fmt.Sprint(err),
		}
	}
	return serializer.ErrorResponse{
		StatusCode: 40001,
		StatusMsg:  "参数错误",
		Error:      fmt.Sprint(err),
	}
}
