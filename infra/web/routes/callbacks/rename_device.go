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

	err = web.GetController().Rename(oldName, body.NewName)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)

		responseError := controllers.ErrorResponse{
			ErrorMessage: err.Error(),
			ErrorCode:    controllers.NetStatInternalError,
		}
		ctx.JSON(responseError)
	}

	ctx.StatusCode(iris.StatusNoContent)
}
