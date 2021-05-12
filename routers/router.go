package routers

import (
	"bee.com/api"
	"bee.com/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	a := r.Group("/api")
	{
		v1 := a.Group("/v1")
		{
			wx := v1.Group("/wx")
			{
				wx.GET("/check_signature", api.WXCheckSignature)
				wx.GET("/get_access_token", api.WxGetAccessTokenFromServer)
			}
		}

		// mock数据
		a.POST("/login", api.MockGetLogin)

		sys := a.Group("/sys")
		{
			sys.GET("/emp", api.MockGetEmpList)
			sys.GET("/roles", api.MockGetRoleList)
		}
		mock := a.Group("/order")
		{
			mock.GET("/list", api.MockGetOrderList)
		}
	}
	return r
}