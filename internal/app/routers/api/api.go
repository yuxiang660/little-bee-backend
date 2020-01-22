package api

import (
	"github.com/yuxiang660/little-bee-server/internal/app/auther"
	"github.com/yuxiang660/little-bee-server/internal/app/controller"
	"github.com/yuxiang660/little-bee-server/internal/app/routers/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterAPI register all APIs to the router.
func RegisterAPI(router *gin.Engine, container *dig.Container) error {
	return container.Invoke(func(
		a auther.Auther,
		loginController controller.ILogin,
	) error {
		api := router.Group("/api")
		{
			api.Use(middleware.UserAuthMiddleware(a,
				middleware.SkipPrefixList("/api/v1/pub/login"),
			))
		}

		v1 := api.Group("/v1")
		{
			pub := v1.Group("/pub")
			{
				gLogin := pub.Group("/login")
				{
					gLogin.POST("", loginController.In)
					gLogin.POST("exit", loginController.Out)
				}
			}
		}

		return nil
	})
}