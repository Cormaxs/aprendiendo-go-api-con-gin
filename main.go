// mi-app-gin/main.go
package main

import (
	"fmt" // Para imprimir mensajes informativos
	"log" // Para registrar errores fatales
	"os"  // Para acceder a variables de entorno

	"github.com/gin-gonic/gin" // El framework web Gin

	// Importa el paquete de rutas. La ruta es relativa al módulo de Go.
	"api-rest/routes" // Asegúrate de que tu proyecto tenga un go.mod (ej. go mod init api-rest)
)

func main() {
	// gin.ReleaseMode es recomendado para producción para deshabilitar mensajes de debug.
	// gin.DebugMode es útil durante el desarrollo.
	gin.SetMode(gin.ReleaseMode)

	// gin.Default() incluye middleware de logging y recuperación de panics por defecto.
	router := gin.Default()

	// Llama a la función de otro paquete para registrar todas las rutas.
	routes.SetupRoutes(router)

	// Intenta obtener el puerto de una variable de entorno llamada "PORT". si no esta definida usa 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Puerto por defecto
	}

	// router.Run() es una llamada bloqueante; si devuelve un error, algo salió mal.
	err := router.Run(":" + port)

	// Manejo de errores: Si el servidor no pudo iniciar, registra el error y termina el programa.
	if err != nil {
		log.Fatalf("El servidor falló al iniciar en el puerto %s: %v", port, err)
	}

	// Este mensaje solo se verá si el servidor arranca y luego se detiene (ej. por una señal de interrupción), ya que router.Run() es bloqueante.
	fmt.Printf("Servidor Gin escuchando en el puerto %s\n", port)
}
