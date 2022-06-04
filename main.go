package main

import (
	_ "github.com/SherzodAbdullajonov/e-wallet/docs"
	"github.com/SherzodAbdullajonov/e-wallet/middleware"
	"github.com/SherzodAbdullajonov/e-wallet/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API document title
// @version version(1.0)
// @description Description of specifications
// @Precautions when using termsOfService specifications

// @contact.name API supporter
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name license(Mandatory)
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /

func main() {
	router := gin.Default()
	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Use(gin.Logger())
	routes.UseRoutes(router)
	router.Use(middleware.Authentication())
	router.Run(":8000")
}
