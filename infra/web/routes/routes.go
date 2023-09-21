package routes

import (
	"bytes"
	"encoding/json"
	"io"
	"network-health/controllers"
	"network-health/infra/web/routes/callbacks"
	"strings"

	iris "github.com/kataras/iris/v12"
	"github.com/samber/lo"
)

type Router struct{}

func (r *Router) Route(i *iris.Application) {
	devicesHealthAPI := i.Party("/v1/network")

	i.Handle("ALL", "/*", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusNotFound)

		responseError := controllers.ErrorResponse{
			ErrorMessage: "Route not found",
			ErrorCode:    controllers.NetStatNotFound,
		}
		ctx.JSON(responseError)
	})

	devicesHealthAPI.Get("/health", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{"status": "ok"})
		return
	})

	devicesHealthAPI.Get("/devices/health", callbacks.GetDevicesHealth)
	devicesHealthAPI.Patch("/devices/rename/:name", CheckBodySupport, callbacks.RenameDevice)

}

func CheckBodySupport(ctx iris.Context) {
	contentType := ctx.GetHeader("Content-Type")

	if !lo.Contains([]string{"application/json"}, contentType) {
		responseError := controllers.ErrorResponse{
			ErrorMessage: "Unsupported body",
			ErrorCode:    controllers.NetStatUnsupportedRequest,
		}
		ctx.JSON(responseError)
		ctx.StatusCode(iris.StatusUnsupportedMediaType)
		return
	}

	body, err := io.ReadAll(ctx.Request().Body)
	if err == nil {
		if !json.Valid(body) || !strings.HasPrefix(string(body), "{") || !strings.HasSuffix(string(body), "}") {
			responseError := controllers.ErrorResponse{
				ErrorMessage: "Unsupported body",
				ErrorCode:    controllers.NetStatUnsupportedRequest,
			}
			ctx.JSON(responseError)
			ctx.StatusCode(iris.StatusUnsupportedMediaType)
			return
		}
	}

	ctx.Request().Body = io.NopCloser(bytes.NewBuffer(body))
	ctx.Next()
}
