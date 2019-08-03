package paycenter

import (
	"errors"
	"log"

	"github.com/fromiuan/epaylinks/lib"
)

type Client struct {
	pfxData  []byte        // pfx文件数据
	password string        // 证书密码
	signNo   string        // 证书序列号
	merchant string        // 商户编号
	dev      bool          // 开发模式
	debug    bool          // debug信息
	encoding *lib.Encoding // 加密解密
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

func (c *Client) doReuest(method, urls string, parms map[string]interface{}, rsp interface{}) error {
	if c.dev {
		urls = testHost + urls
	} else {
		urls = proHost + urls
	}
	if c.debug {
		log.Println("urls：", urls)
	}

	var req *lib.HTTPRequest
	if method == "post" {
		req = lib.Post(urls)
	} else if method == "get" {
		req = lib.Get(urls)
	} else {
		return errors.New("暂时只支持post和get")
	}

	// 修改签名
	parmsSign := lib.ParmsSign(parms)
	sign, err := c.encoding.Sign([]byte(parmsSign))
	if err != nil {
		return err
	}

	parms["sign"] = sign
	if c.debug {
		log.Println("parms:", parms)
	}
	for key, value := range parms {
		req.Param(key, lib.ToString(value))
	}
	err = req.ToXML(rsp)
	if c.debug {
		str, err := req.String()
		log.Println("req str info:", str, err)
	}
	if err != nil {
		return err
	}
	return nil
}
