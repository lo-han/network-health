package web

import "network-health/controllers"

var instantiatedController *controllers.Controller = nil

func SetController(controller *controllers.Controller) {
	instantiatedController = controller
}

func GetController() *controllers.Controller {
	return instantiatedController
}
