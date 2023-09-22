package callbacks

import (
	"network-health/infra/icmp"
	"network-health/infra/web"

	iris "github.com/kataras/iris/v12"
)

func GetDevicesHealth(ctx iris.Context) {
	icmpConnectionHandler := icmp.NewICMPConnectivityHandler()

	response, err := web.Controller().Check(icmpConnectionHandler)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(response)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(response)
}
