package controllers

import (
	"encoding/json"
	"network-health/core/usecases/check"
)

var (
	NetStatOK                 = "NETSTAT_200"
	NetStatNoContent          = "NETSTAT_204"
	NetStatBadRequest         = "NETSTAT_400"
	NetStatNotFound           = "NETSTAT_404"
	NetStatUnsupportedRequest = "NETSTAT_415"
)

type ControllerResponse struct {
	Code    string                 `json:"code"`
	Content map[string]interface{} `json:"content"`
}

func NewControllerResponse(code string, content *check.DeviceStatus) *ControllerResponse {
	bytes, _ := json.Marshal(content)

	m := make(map[string]interface{})
	_ = json.Unmarshal(bytes, &m)

	return &ControllerResponse{
		Content: m,
		Code:    code,
	}
}

func NewControllerEmptyResponse(code string) *ControllerResponse {
	return &ControllerResponse{
		Content: map[string]interface{}{},
		Code:    code,
	}
}

func NewControllerError(code string, message string) *ControllerResponse {
	return &ControllerResponse{
		Content: map[string]interface{}{"error": message},
		Code:    code,
	}
}
