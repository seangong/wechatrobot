package main

import (
	"flag"
	"net/http"

	"wechatrobot/api"
	"wechatrobot/model"

	"github.com/gin-gonic/gin"
)

var (
	h        bool
	RobotKey string
	addr     string
	proxy    string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&proxy, "proxy", "", "http proxy url")
	flag.StringVar(&RobotKey, "RobotKey", "", "wechatrobot token")
	flag.StringVar(&addr, "addr", ":8989", "listen addr")
}

func main() {

	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	client := api.InitClient(proxy)

	// gin 框架
	router := gin.Default()

	// 接收 post 请求，路由为 webhook
	router.POST("/webhook", func(c *gin.Context) {
		var notification model.Notification
		err := c.BindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 获取机器人 token
		RobotKey := c.DefaultQuery("key", RobotKey)

		// 发送请求到机器人接口，支持通过代理访问机器人
		err = api.Send(notification, RobotKey, client)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusOK, gin.H{"message": "send to wechatbot successful!"})
	})
	router.Run(addr)
}
