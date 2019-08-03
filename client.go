package epaylinks

import (
	"encoding/json"
	"log"
	"time"

	"github.com/fromiuan/epaylinks/lib"
)

type Client struct {
	pfxData  []byte // pfx文件数据
	password string // 证书密码
	signNo   string // 证书序列号
	merchant string // 商户编号
	dev      bool   // 开发模式
	debug    bool   // debug信息
	encoding *lib.Encoding
}

func NewClient(pfxData []byte, password, signNo, merchant string) *Client {
	c := &Client{
		pfxData:  pfxData,
		password: password,
		signNo:   signNo,
		merchant: merchant,
		encoding: lib.NewEncding(pfxData, password),
	}
	if signNo == "" {
		c.signNo = c.encoding.GetCertificate()
	}
	return c
}

// set Develop
func (c *Client) SetDev(dev bool) {
	c.dev = dev
}

// set debug
func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) Setting(cli *Client) {
	c = cli
}

func (c *Client) GetSignNo() string {
	return c.signNo
}

// verify
func (c *Client) Verify(sign, body string) error {
	encoding := lib.NewEncding(c.pfxData, c.password)
	return encoding.Verify(sign, body)
}

// do req
func (c *Client) doPostReq(reqUrl string, params, rsp interface{}) error {
	if c.dev {
		reqUrl = testHost + reqUrl
	} else {
		reqUrl = proHost + reqUrl
	}

	req := lib.Post(reqUrl)
	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return err
	}
	signStr, err := c.encoding.Sign(paramsBytes)
	if err != nil {
		return err
	}
	req.Header("x-efps-sign-no", c.GetSignNo())
	req.Header("x-efps-sign-type", "SHA256withRSA")
	req.Header("x-efps-sign", signStr)
	req.Header("x-efps-timestamp", time.Now().Format("20060102150405"))
	req.Header("Content-Type", "application/json")
	req.Body(paramsBytes)
	err = req.ToJSON(rsp)
	if c.debug {
		str, err := req.String()
		log.Println("req str info:", str, err)
	}
	if err != nil {
		return err
	}
	return err
}
