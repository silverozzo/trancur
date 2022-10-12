package http

import (
	"trancur/http/controller"
)

type Config interface {
	controller.Config

	GetSelfHttpPort() string
}
