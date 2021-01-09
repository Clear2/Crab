package routers

import (
	"bee.com/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

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
	}
	return r
}