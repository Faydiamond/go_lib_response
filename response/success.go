package response

import (
	"github.com/Faydiamond/microservice_meta/meta"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Meta    *meta.Meta  `json:"meta,omitempty"`
}

func success(msg string, data interface{}, meta *meta.Meta, code int) Response {
	return &SuccessResponse{
		Message: msg,
		Status:  code,
		Data:    data,
		Meta:    meta,
	}
}
