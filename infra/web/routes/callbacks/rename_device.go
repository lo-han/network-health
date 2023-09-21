package callbacks

import (
	"encoding/json"
	"network-health/controllers"
	"network-health/infra/web"

	iris "github.com/kataras/iris/v12"
)

type RenameBody struct {
	NewName string `json:"new_name"`
}

func RenameDevice(ctx iris.Context) {
	var body RenameBody

	oldName := ctx.Params().Get("name")

	bodyBytes, _ := ctx.GetBody()

	err := json.Unmarshal(bodyBytes, &body)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)

		responseError := controllers.ErrorResponse{
			ErrorMessage: "Bad request",
			ErrorCode:    controllers.NetStatBadRequest,
		}
		ctx.JSON(responseError)
		return
	}

	if body.NewName == "" {
		ctx.StatusCode(iris.StatusBadRequest)

		responseError := controllers.ErrorResponse{
			ErrorMessage: "Bad request",
			ErrorCode:    controllers.NetStatBadRequest,
		}
		ctx.JSON(responseError)
		return
	}

	err = web.GetController().Rename(oldName, body.NewName)

	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)

		responseError := controllers.ErrorResponse{
			ErrorMessage: err.Error(),
			ErrorCode:    controllers.NetStatNotFound,
		}
		ctx.JSON(responseError)
		return
	}

	ctx.StatusCode(iris.StatusNoContent)
}
