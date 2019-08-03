package paycenter

import (
	"encoding/xml"
	"errors"

	"github.com/fromiuan/epaylinks/lib"
)

type GetOI struct {
	Version         string `json:"version"`          // 接口版本
	SignType        string `json:"sign_type"`        // 签名类型
	CertId          string `json:"certId"`           // 证书序列号
	Sign            string `json:"sign"`             // 签名
	Partner         string `json:"partner"`          // 商户号
	OutTradeNo      string `json:"out_trade_no"`     // 商户系统订单号
	TotalFee        string `json:"total_fee"`        // 订单金额
	CurrencyType    string `json:"currency_type"`    // 货币代码
	ReturnUrl       string `json:"return_url"`       // 前台返回URL
	NotifyUrl       string `json:"notify_url"`       // 后台通知URL
	OrderCreateIp   string `json:"order_create_ip"`  // 订单创建IP
	PayId           string `json:"pay_id"`           // 银行直连参数
	Base64Memo      string `json:"base64_memo"`      // 商品名称
	StoreOiType     string `json:"store_oi_type"`    // 网关类型
	TimeOut         string `json:"time_out"`         // 订单超时时间
	MerchantAccount string `json:"merchant_account"` // 银行卡卡号
	CardType        string `json:"card_type"`        // 支持卡类型
}

type GetIORsp struct {
	XMLName    xml.Name `xml:"root"`
	Version    string   `xml:"Version"`      // 接口版本
	SignType   string   `xml:"sign_type"`    // 签名类型
	Sign       string   `xml:"Sign"`         // 签名
	Partner    string   `xml:"Partner"`      // 商户号
	OutTradeNo string   `xml:"out_trade_no"` // 商户系统订单号
	PayNo      string   `xml:"pay_no"`       // 支付单号
	Amount     string   `xml:"amount"`       // 订单金额
	PayResult  string   `xml:"pay_result"`   // 支付结果
	PayTime    string   `xml:"pay_time"`     // 支付时间
	SettDate   string   `xml:"sett_date"`    // 清算日期
	SettTime   string   `xml:"sett_time"`    // 清算时间
	Base64Memo string   `xml:"base64_memo"`  // 商品名称
	ChannelId  string   `xml:"channel_id"`   // 支付渠道
	DecFeeRate string   `xml:"dec_fee_rate"` // 费率
	DecFee     string   `xml:"dec_fee"`      // 手续费
}

func (c *Client) GetOI(goi *GetOI) (rsp *GetIORsp, err error) {
	rsp = new(GetIORsp)

	err = goi.checkParms()
	if err != nil {
		return rsp, err
	}
	mp, err := lib.ToMap(goi)
	if err != nil {
		return rsp, err
	}
	err = c.doReuest("post", getoi, mp, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

func (goi *GetOI) checkParms() error {
	if goi.Version == "" {
		return errors.New("接口版本不能为空")
	}
	if goi.SignType != "SHA256withRSA" {
		return errors.New("请输入正确的签名类型")
	}
	if goi.CertId == "" {
		return errors.New("证书序列号不能为空")
	}
	if goi.Partner == "" {
		return errors.New("商户号不能为空")
	}
	if goi.OutTradeNo == "" {
		return errors.New("商户系统订单号不能为空")
	}
	if goi.TotalFee == "" {
		return errors.New("订单金额不能为空")
	}
	if goi.CurrencyType == "" {
		return errors.New("货币代码不能为空")
	}
	if goi.ReturnUrl == "" {
		return errors.New("前台返回URL不能为空")
	}
	return nil
}
