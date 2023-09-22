package controllers

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

func NewControllerResponse(code string, content map[string]interface{}) *ControllerResponse {
	return &ControllerResponse{
		Content: content,
		Code:    code,
	}
}

func NewControllerError(code string, message string) *ControllerResponse {
	return &ControllerResponse{
		Content: map[string]interface{}{"error": message},
		Code:    code,
	}
}
