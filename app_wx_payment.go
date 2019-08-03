package epaylinks

import (
	"errors"
)

type AppWxPayment struct {
	OutTradeNo           string     `json:"outTradeNo"`           // 商户订单号
	CustomerCode         string     `json:"customerCode"`         // 商户号
	ClientIp             string     `json:"clientIp"`             // 客户端ip
	OrderInfo            *OrderInfo `json:"orderInfo"`            // 商品订单信息
	PayAmount            int64      `json:"payAmount"`            // 支付金额
	PayCurrency          string     `json:"payCurrency"`          // 支付币种
	ChannelType          string     `json:"channelType"`          // 渠道类型
	NotifyUrl            string     `json:"notifyUrl"`            // 支付结果异步通知地址
	AttachData           string     `json:"attachData"`           // 附加数据
	TransactionStartTime string     `json:"transactionStartTime"` // 交易开始时间
	TransactionEndTime   string     `json:"transactionEndTime"`   // 交易超时时间
	NeedSplit            bool       `json:"needSplit"`            // 是否分账
	NoCreditCards        bool       `json:"noCreditCards"`        // 是否禁止信用卡支付
	NonceStr             string     `json:"nonceStr"`             // 随机字符串
}

type AppWxPaymentRsp struct {
	ReturnCode string      `json:"returnCode"` // 返回状态码
	ReturnMsg  string      `json:"returnMsg"`  // 返回信息
	AppPayInfo *AppPayInfo `json:"appPayInfo"` // 商户APP调起微信APP支付所需要的参数对象Json
	OutTradeNo string      `json:"outTradeNo"` // 商户订单号
	Amount     string      `json:"amount"`     // 支付金额
	NonceStr   string      `json:"nonceStr"`   // 随机字符串

}

type AppPayInfo struct {
	AppId     string `json:"appId"`     // 商户APPID
	NonceStr  string `json:"nonceStr"`  // 随机字符串
	Package   string `json:"Package"`   // package
	PartnerId string `json:"partnerId"` // partnerid
	PaySign   string `json:"paySign"`   // 签名
	PrepayId  string `json:"prepayId"`  // prepayid
	TimeStamp string `json:"timeStamp"` // 时间戳
}

// 微信APP支付接口
func (c *Client) AppWxPayment(awp *AppWxPayment) (rsp *AppWxPaymentRsp, err error) {
	rsp = new(AppWxPaymentRsp)

	err = awp.checkParms()
	if err != nil {
		return rsp, err
	}

	err = c.doPostReq(appWxPayMent, awp, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (awp *AppWxPayment) checkParms() error {
	if awp.OutTradeNo == "" {
		return errors.New("商户订单号不能为空")
	}
	if awp.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if awp.ClientIp == "" {
		return errors.New("用户终端IP不能为空")
	}
	if awp.OrderInfo == nil {
		return errors.New("商品订单信息不能为空")
	}
	if awp.PayAmount <= 0 {
		return errors.New("支付金额小于等于0")
	}
	if awp.PayCurrency == "" {
		return errors.New("支付币种不能为空")
	}
	if awp.ChannelType != "01" && awp.ChannelType != "02" {
		return errors.New("渠道类型必须为空01或者02")
	}
	if awp.NotifyUrl == "" {
		return errors.New("支付结果通知地址为空")
	}
	if awp.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	if awp.TransactionStartTime == "" {
		return errors.New("交易开始时间必填")
	}
	return nil
}
