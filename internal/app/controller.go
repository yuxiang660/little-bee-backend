package app

import (
	"github.com/yuxiang660/little-bee-server/internal/app/controller/login"
	"github.com/yuxiang660/little-bee-server/internal/app/controller/user"
	"go.uber.org/dig"
)

// InjectController injects an controller constructor to dig container.
func InjectController(container *dig.Container) func() {

	_ = container.Provide(login.New)
	_ = container.Provide(user.New)

	return nil
}
