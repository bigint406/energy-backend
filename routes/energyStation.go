package routes

import (
	"energy/api/pageDataPresent"

	"github.com/gin-gonic/gin"
)

func EnergyStationRouter(router *gin.RouterGroup) {
	router.GET("basicdata", pageDataPresent.BasicData)
	router.GET("basicdatalist", pageDataPresent.BasicDataList)
}
