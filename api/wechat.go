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
	const url = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx5c384a6e1d773bab&secret=3b83dd39cc7e3257ab03d8c5c6774ef9"
	//url := fmt.Sprintf("%s/api/10008/securityUser/find", baseURL)
	rs, err := utils.New(url, "POST").Request()
	if err != nil {
		log.Fatal(err)
	}
	utils.ResSuccess(c, rs)
	log.Println("----->>>公众号校验失败")

}

func WechatMessageReceive(c *gin.Context) {
	//var textMsg models.WeTextMsg
	//err := c.ShouldBindXML(&textMsg)
	//if err != nil {
	//	log.Println("------->>> XML数据包解析失败")
	//}

	//log.Println("收到消息--->>", textMsg.Content)
}