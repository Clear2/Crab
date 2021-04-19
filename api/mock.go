package api

import (
	"bee.com/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetOrderList(c * gin.Context)  {
	data := map[string]string{
		"password": "HUAxia123",
		"username": "5628",
	}

	resp, err := utils.New("https://pzgadmin.pzgmall.com/api/login", "POST").WithRawBody(data).Request()

	rs, err := utils.New("https://pzgadmin.pzgmall.com/api/8001/order/find", "POST").Request()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	c.JSON(http.StatusOK, gin.H{"data": rs})
}

/*
"password": {

		},
		"username": {
			"5628",
		},*/