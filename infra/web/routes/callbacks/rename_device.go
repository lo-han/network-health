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

		responseError := controllers.NewControllerError(controllers.NetStatBadRequest, "Bad request. Wrong type parameter")
		ctx.JSON(responseError)
		return
	}

	response, err := web.Controller().Rename(oldName, body.NewName)

	if err != nil {
		switch response.Code {
		case controllers.NetStatBadRequest:
			ctx.StatusCode(iris.StatusBadRequest)
		case controllers.NetStatNotFound:
			ctx.StatusCode(iris.StatusNotFound)
		}

		ctx.JSON(response)
		return
	}

	ctx.StatusCode(iris.StatusNoContent)
}
