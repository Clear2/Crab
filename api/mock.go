package api

import (
	"bee.com/utils"
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
)

var baseURL string= "http://172.16.7.3:20001"

func MockGetLogin(c *gin.Context)  {
	//data := map[string]string{
	//	"password": "HUAxia123",
	//	"username": "5628",
	//}
	//resp, err := utils.New("https://pzgadmin.pzgmall.com/api/login", "POST").WithRawBody(data).Request()
	data := map[string]string{
		"password": "pzg@123456",
		"username": "00000777",
	}
	url := fmt.Sprintf("%s/api/login", baseURL)
	resp, err := utils.New(url, "POST").WithRawBody(data).Request()
	if err != nil {
		log.Fatal(err)
	}
	utils.ResSuccess(c, resp)
}
func MockGetOrderList(c * gin.Context)  {
	url := fmt.Sprintf("%s/api/8001/order/find", baseURL)
	rs, err := utils.New(url, "POST").Request()
	if err != nil {
		log.Fatal(err)
	}
	utils.ResSuccess(c, rs)
}

func MockGetEmpList(c *gin.Context) {
	url := fmt.Sprintf("%s/api/10008/securityUser/find", baseURL)
	rs, err := utils.New(url, "POST").Request()
	if err != nil {
		log.Fatal(err)
	}
	utils.ResSuccess(c, rs)
}

func MockGetRoleList(c *gin.Context) {
	url := fmt.Sprintf("%s/api/10009/security/findRole", baseURL)
	rs, err := utils.New(url, "POST").Request()
	if err != nil {
		log.Fatal(err)
	}
	utils.ResSuccess(c, rs)
}


// pzg@123456
func MockGetPmsList (c * gin.Context) {
	url := fmt.Sprintf("%s/api/8001/order/find", baseURL)
	rs, err := utils.New(url, "POST").Request()
	if err != nil {
		log.Fatal(err)
	}
	utils.ResSuccess(c, rs)
}