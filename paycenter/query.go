package paycenter

import (
	"encoding/xml"
	"errors"

	"github.com/fromiuan/epaylinks/lib"
)

type Gateways struct {
	Version    string // 接口版本
	SignType   string // 签名类型
	CertId     string // 证书序列号
	Sign       string // 签名
	Partner    string // 商户号
	TransType  string // 交易类型
	OutTradeNo string // 商户系统订单号
	PayNo      string // 易票联支付号

}

type GatewaysRsp struct {
	XMLName         xml.Name `xml:"root"`
	Version         string   `xml:"version"`         // 接口版本
	SignType        string   `xml:"sign_type"`       // 签名类型
	sign            string   `xml:"sign"`            // 签名
	RespCode        string   `xml:"resp_code"`       // 响应码
	RespDesc        string   `xml:"resp_desc"`       // 响应描述
	Partner         string   `xml:"partner"`         // 商户号
	OutTradeNo      string   `xml:"out_trade_no"`    // 商户系统订单号
	CurrCode        string   `xml:"curr_code"`       // 货币代码
	PayNo           string   `xml:"pay_no"`          // 支付单号
	Amount          string   `xml:"amount"`          // 订单金额
	PayResult       string   `xml:"pay_result"`      // 支付结果
	PayTime         string   `xml:"pay_time"`        // 支付时间
	SettDate        string   `xml:"sett_date"`       // 清算日期
	SettTime        string   `xml:"sett_time"`       // 清算时间
	ChannelId       string   `xml:"channel_id"`      // 支付渠道
	DecFeeRate      string   `xml:"dec_fee_rate"`    // 费率
	DecFee          string   `xml:"dec_fee"`         // 手续费
	ChannelSerialNo string   `xml:"channelSerialNo"` // 上游订单号
}

func (c *Client) Query(g *Gateways) (rsp *GatewaysRsp, err error) {
	rsp = new(GatewaysRsp)

	err = g.checkParms()
	if err != nil {
		return rsp, err
	}
	mp, err := lib.ToMap(g)
	if err != nil {
		return rsp, err
	}
	err = c.doReuest("post", getoi, mp, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

func (g *Gateways) checkParms() error {
	if g.Version == "" {
		return errors.New("接口版本不能为空")
	}
	if g.SignType != "SHA256withRSA" {
		return errors.New("请输入正确的签名类型")
	}
	if g.CertId == "" {
		return errors.New("证书序列号不能为空")
	}
	if g.Partner == "" {
		return errors.New("商户号不能为空")
	}
	if g.TransType == "" {
		return errors.New("订单金额不能为空")
	}
	if g.OutTradeNo == "" && g.PayNo == "" {
		return errors.New("三方单号和易票联单号任意一个")
	}
	return nil
}
