package controllers

import (
	"github.com/gin-gonic/gin" // Importa Gin para poder usar gin.Context
)

// PingHandler es la función controladora para la ruta "/ping".
// Recibe un *gin.Context, que contiene toda la información de la solicitud HTTP y métodos para construir la respuesta.
func PingHandler(c *gin.Context) {
	// gin.H es un atajo para map[string]interface{}, utilizado para crear objetos JSON de forma concisa.
	// Al centralizar la lógica de respuesta aquí, las rutas solo necesitan "apuntar" a esta función.
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
