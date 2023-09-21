package callbacks

import (
	"network-health/controllers"
	"network-health/infra/icmp"
	"network-health/infra/web"

	iris "github.com/kataras/iris/v12"
)

func GetDevicesHealth(ctx iris.Context) {
	icmpConnectionHandler := icmp.NewICMPConnectivityHandler()

	response, err := web.GetController().Check(icmpConnectionHandler)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)

		responseError := controllers.ErrorResponse{
			ErrorMessage: err.Error(),
			ErrorCode:    controllers.NetStatInternalError,
		}
		ctx.JSON(responseError)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(response)
}
