package api

import (
	"bee.com/utils"
	"github.com/gin-gonic/gin"
	"log"
)

const Token = "Clear"

func WXCheckSignature(c *gin.Context)  {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	ok := utils.CheckSignature(signature, timestamp, nonce, Token)

	if !ok {
		log.Println("----->>>公众号校验失败")
		return
	}
	log.Println("----->>公众号校验成功")
	c.Writer.WriteString(echostr)
}

func WxGetAccessTokenFromServer(c *gin.Context) {
}

func WechatMessageReceive(c *gin.Context) {
	//var textMsg models.WeTextMsg
	//err := c.ShouldBindXML(&textMsg)
	//if err != nil {
	//	log.Println("------->>> XML数据包解析失败")
	//}

	//log.Println("收到消息--->>", textMsg.Content)
}