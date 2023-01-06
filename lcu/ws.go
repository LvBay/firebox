package lcu

import (
	"crypto/tls"
	"fmt"
	"github.com/gogf/gf/v2/net/gclient"
	"log"
	"time"
)

func (c *Client) SubscribeGamePhase(url string) {
	client := gclient.NewWebSocket()
	client.HandshakeTimeout = time.Second  // 设置超时时间
	client.TLSClientConfig = &tls.Config{} // 设置 tls 配置

	conn, _, err := client.Dial("ws://127.0.0.1:8001/ws", nil)
	if err != nil {
		log.Println("dial ws failed", err)
		panic(err)
	}
	c.WsClient = client
	c.WsConn = conn

	go func() {
		for {
			mt, data, err := c.WsConn.ReadMessage()
			if err != nil {
				log.Println("read ws failed", err)
				return
			}
			fmt.Println(mt, data)
			c.GamePhaseCh <- string(data)
		}
	}()
}
