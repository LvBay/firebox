package lcu

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"os/exec"
	"strings"
)

type AuthInfo struct {
	Token   string
	AppID   string
	AppPort string
}

type Client struct {
	AuthInfo    AuthInfo
	Client      *http.Client
	Protocol    string
	WsClient    *gclient.WebSocketClient
	WsConn      *websocket.Conn
	GamePhaseCh chan string
}

func NewClient() (*Client, error) {
	auth, err := GetAuthInfo()
	if err != nil {
		return nil, err
	}
	tr := http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	ck, _ := cookiejar.New(nil)
	cli := http.Client{Transport: &tr, Jar: ck}
	c := &Client{Client: &cli, AuthInfo: *auth}
	if isMockLcu() {
		c.Protocol = "http"
		c.Client = http.DefaultClient
		// c.SubscribeGamePhase("")
		// c.GamePhaseCh = make(chan string)
	}
	return c, nil
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	bs, err := io.ReadAll(resp.Body)
	return bs, nil
}

func (c *Client) buildRequest(method, path string, body interface{}) *http.Request {
	b := bytes.NewReader(nil)
	if body != nil {
		bs, _ := json.Marshal(body)
		b = bytes.NewReader(bs)
	}

	protocol := "https"
	if c.Protocol != "" {
		protocol = c.Protocol
	}
	req, _ := http.NewRequest(method,
		fmt.Sprintf("%s://127.0.0.1:%s%s", protocol, c.AuthInfo.AppPort, path), b)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(c.AuthInfo.Token))))

	return req
}

func (c *Client) Do(method, path string, body interface{}) ([]byte, error) {
	req := c.buildRequest(method, path, body)

	bs, err := c.do(req)
	// log.Println("debug info method", req.Method)
	// log.Println("debug info url", req.URL)
	// log.Println("debug info resp", string(bs))
	if err != nil {
		log.Println("http do failed", "err", err, "path", path)
	}
	return bs, err
}

func isMockLcu() bool {
	m := os.Getenv("LCU_MOCK")
	if m == "true" {
		return true
	}
	return false
}

var isMock bool

func init() {
	isMock = isMockLcu()
}

func GetAuthInfo() (*AuthInfo, error) {
	if isMock {
		return &AuthInfo{AppPort: "8001"}, nil
	}
	cmd := exec.Command("cmd", "/c", "wmic PROCESS WHERE name='LeagueClientUx.exe'")
	out, err := cmd.Output()
	fmt.Println(string(out), err)
	if len(out) == 0 {
		return nil, errors.New("游戏未启动")
	}
	pInfo := string(out)
	token := getKey(pInfo, "--remoting-auth-token=")
	port := getKey(pInfo, "--app-port=")
	pid := getKey(pInfo, "--app-pid=")
	return &AuthInfo{
		Token:   "riot:" + token,
		AppID:   pid,
		AppPort: port,
	}, nil
}

func getKey(out string, key string) string {
	list := strings.Split(out, key)
	list = strings.Split(list[1], `"`)
	return list[0]
}
