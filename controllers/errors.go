package controllers

var (
	NetStatBadRequest    = "NETSTAT_400"
	NetStatNotFound      = "NETSTAT_404"
	NetStatRuleError     = "NETSTAT_422"
	NetStatInternalError = "NETSTAT_500"
)

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
	ErrorCode    string `json:"code"`
}