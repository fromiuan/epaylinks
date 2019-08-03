package epaylinks

import (
	"errors"
)

type AliJSAPIPayment struct {
	OutTradeNo           string     `json:"outTradeNo"`           // 商户订单号
	BuyerId              string     `json:"buyerId"`              // 买家的支付宝唯一用户号
	BuyerLogonId         string     `json:"buyerLogonId"`         // 买家的支付宝账号
	CustomerCode         string     `json:"customerCode"`         // 商户号
	OrderInfo            *OrderInfo `json:"orderInfo"`            // 商品订单信息
	PayAmount            int64      `json:"payAmount"`            // 支付金额
	PayCurrency          string     `json:"payCurrency"`          // 支付币种
	ChannelType          string     `json:"channelType"`          // 渠道类型
	NoCreditCards        bool       `json:"noCreditCards"`        // 是否禁止信用卡支付
	NotifyUrl            string     `json:"notifyUrl"`            // 支付结果通知地址
	AttachData           string     `json:"attachData"`           // 附加数据
	TransactionStartTime string     `json:"transactionStartTime"` // 交易开始时间
	TransactionEndTime   string     `json:"transactionEndTime"`   // 交易超时时间
	NonceStr             string     `json:"nonceStr"`             // 随机字符串
	NeedSplit            bool       `json:"needSplit"`            // 是否分账
	RechargeMemCustCode  string     `json:"rechargeMemCustCode"`  // 充值会员客户编码
}

type AliJSAPIPaymentRsp struct {
	ReturnCode    string `json:"returnCode"`    // 返回状态码
	ReturnMsg     string `json:"returnMsg"`     // 返回信息
	OutTradeNo    string `json:"outTradeNo"`    // 商户订单号
	Amount        int64  `json:"amount"`        // 支付金额
	NonceStr      string `json:"nonceStr"`      // 随机字符串
	AlipayTradeNo string `json:"alipayTradeNo"` // 支付宝交易号
}

// 支付宝服务窗支付接口
func (c *Client) AliJSAPIPayment(ajsapip *AliJSAPIPayment) (rsp *AliJSAPIPaymentRsp, err error) {
	rsp = new(AliJSAPIPaymentRsp)

	err = ajsapip.checkParms()
	if err != nil {
		return rsp, err
	}

	err = c.doPostReq(aliJSAPIPayment, ajsapip, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (ajsapip *AliJSAPIPayment) checkParms() error {
	if ajsapip.OutTradeNo == "" {
		return errors.New("商户订单号不能为空")
	}
	if ajsapip.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if ajsapip.OrderInfo == nil {
		return errors.New("商品订单信息不能为空")
	}
	if ajsapip.PayAmount <= 0 {
		return errors.New("支付金额小于等于0")
	}
	if ajsapip.PayCurrency == "" {
		return errors.New("支付币种不能为空")
	}
	if ajsapip.ChannelType != "01" && ajsapip.ChannelType != "02" {
		return errors.New("渠道类型必须为空01或者02")
	}
	if ajsapip.NotifyUrl == "" {
		return errors.New("支付结果通知地址为空")
	}
	if ajsapip.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	if ajsapip.TransactionStartTime == "" {
		return errors.New("交易开始时间必填")
	}
	return nil
}
