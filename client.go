package epaylinks

import (
	"github.com/fromiuan/epaylinks/lib"
)

type Client struct {
	pfxData  []byte // pfx文件数据
	password string // 证书密码
	signNo   string // 证书序列号
	merchant string // 商户编号
	dev      bool   // 开发模式
	debug    bool   // debug信息
}

func NewClient(pfxData []byte, password, signNo, merchant string) *Client {
	return &Client{
		pfxData:  pfxData,
		password: password,
		signNo:   signNo,
		merchant: merchant,
	}
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

// verify
func (c *Client) Verify(sign, body string) error {
	encoding := lib.NewEncding(c.pfxData, c.password)
	return encoding.Verify(sign, body)
}
