// mi-app-gin/routes/routes.go
package routes

import (
	"api-rest/controllers"

	"github.com/gin-gonic/gin" // Importa Gin para poder usar gin.Engine y gin.Context
)

// SetupRoutes es una función pública que recibe un puntero al router de Gin.
// Su propósito es registrar todas las rutas de la aplicación en el router proporcionado.
// Esto permite que el archivo main.go no tenga que conocer los detalles de cada ruta.
func SetupRoutes(router *gin.Engine) {
	// Aquí se asocia la URL "/ping" con una función handler.
	router.GET("/ping", controllers.PingHandler)

	// O agrupar rutas para una mejor organización (útil para APIs más grandes):
	api := router.Group("/api")
	{
		api.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "todos los usuarios aquí"})
		})

		api.GET("/personaje", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "personaje elegido"})
		})
	}

}
