package main

import (
	"enroll/database"
	"enroll/routes"
	"enroll/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Main struct {
	appRouting *gin.Engine
}

func (mainModule *Main) initServer() error {
	var err error

	err = utils.LoadConfig("local")
	if err != nil {
		return err
	}

	mainModule.appRouting = gin.Default()

	err = database.Database.InitConnection()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	mainModule := Main{}

	if mainModule.initServer() != nil {
		return
	}

	v1 := mainModule.appRouting.Group("/api/v1")
	{
		userGroup := v1.Group("/users")
		{
			userRoutes := routes.User{}
			userRoutes.Handler(userGroup)
		}
		authGroup := v1.Group("/auth")
		{
			authRoutes := routes.Auth{}
			authRoutes.Handler(authGroup)

		}
	}

	mainModule.appRouting.Run(utils.ConfigurationEnvs.AppPort)
}
