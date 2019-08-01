package epaylinks

import (
	"errors"
)

type Unifieder struct {
	OutTradeNo           string     `json:"outTradeNo"`           // 商户订单号
	CustomerCode         string     `json:"customerCode"`         // 商户号
	TerminalNo           string     `json:"terminalNo"`           // 终端号
	ClientIp             string     `json:"clientIp"`             // 用户终端 IP
	OrderInfo            *OrderInfo `json:"orderInfo"`            // 商品订单信息
	PayAmount            int64      `json:"payAmount"`            // 支付金额
	PayCurrency          string     `json:"payCurrency"`          // 支付币种(CNY)
	ChannelType          string     `json:"channelType"`          // 渠道类型
	NotifyUrl            string     `json:"notifyUrl"`            // 支付结果通知地址
	RedirectUrl          string     `json:"redirectUrl"`          // 商家支付结果展示地址
	AttachData           string     `json:"attachData"`           // 附加数据
	TransactionStartTime string     `json:"transactionStartTime"` // 随机字符串
	TransactionEndTime   string     `json:"transactionEndTime"`   // 交易开始时间
	NeedSplit            bool       `json:"needSplit"`            // 是否分账
	NoCreditCards        bool       `json:"noCreditCards"`        // 是否禁止信用卡
	NonceStr             string     `json:"nonceStr"`             //随机数
}

type UnifiederRsp struct {
	ReturnCode string `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
	CasherUrl  string `json:"casherUrl"`
	OutTrdeNo  string `json:"outTrdeNo"`
	NonceStr   string `json:"nonceStr"`
}

// 统一下单
func (c *Client) Payment(u *Unifieder) (rsp *UnifiederRsp, err error) {
	rsp = new(UnifiederRsp)

	err = u.checkParms()
	if err != nil {
		return rsp, err
	}

	req := newSetting(unified, c)
	err = req.doPostReq(u, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

// 判断错误
func (u *Unifieder) checkParms() error {
	if u.OutTradeNo == "" {
		return errors.New("商户订单号不能为空")
	}
	if u.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if u.ClientIp == "" {
		return errors.New("用户终端IP不能为空")
	}
	if u.OrderInfo == nil {
		return errors.New("商品订单信息不能为空")
	}
	if u.PayAmount <= 0 {
		return errors.New("支付金额小于等于0")
	}
	if u.PayCurrency == "" {
		return errors.New("支付币种不能为空")
	}
	if u.ChannelType != "01" && u.ChannelType != "02" {
		return errors.New("渠道类型必须为空01或者02")
	}
	if u.NotifyUrl == "" {
		return errors.New("支付结果通知地址为空")
	}
	if u.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	if u.TransactionStartTime == "" {
		return errors.New("交易开始时间必填")
	}
	return nil
}
