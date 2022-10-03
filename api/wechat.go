package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"wechatrobot/model"
	"wechatrobot/pkg"
)

// Send send markdown message to wechatrobot
func Send(notification model.Notification, defaultRobot string, proxy string) (err error) {
	var (
		wechatRobotURL string
		client         *http.Client
	)

	// 获取 markdown 消息结构体 和 机器人 url
	markdown, robotURL, err := pkg.Markdown(notification)

	if err != nil {
		return
	}
	// 序列化 markdown 消息结构为 json 字符串
	data, err := json.Marshal(markdown)
	if err != nil {
		return
	}

	// 是否使用默认机器人
	if robotURL != "" {
		wechatRobotURL = robotURL
	} else {
		wechatRobotURL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + defaultRobot
	}

	// 构造请求结构体
	req, err := http.NewRequest(
		"POST",
		wechatRobotURL,
		bytes.NewBuffer(data))

	if err != nil {
		return
	}

	// 设置 Content-Type 请求头
	req.Header.Set("Content-Type", "application/json")

	// 是否使用代理
	if proxy != "" {
		proxyURL, _ := url.Parse(proxy)

		// 实例化使用代理的 http 客户端
		client = &http.Client{
			Transport: &http.Transport{
				Proxy:           http.ProxyURL(proxyURL),
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 跳过 https 不安全验证
			},
		}
	} else {
		// 实例化 http 客户端
		client = &http.Client{}
	}

	// post 请求机器接口
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	// 关闭响应体
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return
}
