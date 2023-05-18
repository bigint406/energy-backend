package routes

import (
	"energy/api/basicMap"
	"github.com/gin-gonic/gin"
)

func BasicMapRouter(router *gin.RouterGroup) {
	router.GET("basicMap/getAtmosphere", basicMap.GetAtmosphere)
	router.GET("basicMap/getKekong", basicMap.GetKekong)
}
