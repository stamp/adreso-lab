//go:generate bash -c "go get github.com/swaggo/swag/cmd/swag && swag init"

package main

import (
	"github.com/stamp/adreso-lab/api"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	_ "github.com/stamp/adreso-lab/docs"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// @title adreso
// @version 1.1

// @BasePath /v1

func main() {
	logrus.WithFields(logrus.Fields{
		"version": version,
		"commit":  commit,
		"date":    date,
	}).Infof("Startup")

	r := gin.New()
	r.Use(cors.Default())

	v1 := r.Group("/v1")
	{
		addresses := v1.Group("/addresses")
		{
			addresses.GET("", api.ListAddresses)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()

}
