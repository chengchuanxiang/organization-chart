package routes

import (
	"github.com/gin-gonic/gin"
	"oyccx/api/v1"
	"oyccx/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	{
		routerV1.POST("excel/toLead",v1.ToLead1)
		routerV1.GET("excel/getJSON",v1.GetJSON)
	}


	_ = r.Run(utils.HttpPort)
}
