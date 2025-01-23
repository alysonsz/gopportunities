package router

import "github.com/gin-gonic/gin"

func Router() { //O nome da função importa, pois se fosse router(), ela seria uma função local, e não seria possível exportar
	router := gin.Default()
	initializeRoutes(router)
	router.Run(":3030")
}
