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

	devicesHealthAPI.Use(ContentType("PATCH", "application/json"))
	devicesHealthAPI.Use(CheckJSONBody("PATCH"))

	i.Handle("ALL", "/*", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusNotFound)

		responseError := controllers.ErrorResponse{
			ErrorMessage: "Route not found",
			ErrorCode:    controllers.NetStatNotFound,
		}
		ctx.JSON(responseError)
	})

	i.Get("/health", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{"status": "ok"})
	})

	devicesHealthAPI.Get("/devices/health", callbacks.GetDevicesHealth)
	devicesHealthAPI.Patch("/devices/rename/:name", callbacks.RenameDevice)

}

func ContentType(method string, allowedContentTypes ...string) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		contentType := ctx.GetHeader("Content-Type")

		if lo.Contains(allowedContentTypes, contentType) {
			ctx.Next()
			return
		}

		ctx.StatusCode(iris.StatusUnsupportedMediaType)
	}
}

func CheckJSONBody(method ...string) func(ctx iris.Context) {
	methodTypes := make(map[string]bool, len(method))
	for _, m := range method {
		methodTypes[m] = true
	}

	return func(ctx iris.Context) {
		if !methodTypes[ctx.Method()] {
			ctx.Next()
			return
		}

		body, err := io.ReadAll(ctx.Request().Body)
		if err == nil {
			if json.Valid(body) && strings.HasPrefix(string(body), "{") && strings.HasSuffix(string(body), "}") {
				ctx.Request().Body = io.NopCloser(bytes.NewBuffer(body))
				ctx.Next()
				return
			}
		}

		ctx.StatusCode(iris.StatusUnsupportedMediaType)
	}
}
