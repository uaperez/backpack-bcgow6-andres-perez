package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	/*
		gin.Default() -> engine de gin, instancia los middlewares Logger y Recovery de manera instantánea
		es por ello que en la terminal cuando lista las routes disponible, aparecerán siempre 2 por defecto
		más los middlewares que se invoquen en las funciones anónimas del router
	*/
	router := gin.Default()
	// router.Use() -> para utilizar middlewares de autentificación, por ejemplo
	router.GET("/welcome", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name", "Bootcamper")
		// Para el HTTP Code, se puede usar el numero como tal o la constante
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hola " + name,
		})
	})

	router.Run() // para indicar el puerto se le pasa como parametro === router.Run(":5040"). Por defecto corre en el 8080
}
