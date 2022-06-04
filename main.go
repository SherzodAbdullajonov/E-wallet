package main

import (
	"github.com/SherzodAbdullajonov/e-wallet/middleware"
	"github.com/SherzodAbdullajonov/e-wallet/routes"
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	router.Use(gin.Logger())
	routes.UseRoutes(router)
	router.Use(middleware.Authentication())
	router.Run(":8000")
}