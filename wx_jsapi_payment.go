package epaylinks

import (
	"errors"
)

type WxJSAPIPayment struct {
	OutTradeNo           string     `json:"outTradeNo"`           // 商户订单号
	CustomerCode         string     `json:"customerCode"`         // 商户号
	ClientIp             string     `json:"clientIp"`             // 用户的IP
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
	AppId                string     `json:"appId"`                // 微信公众号的APPID
	OpenId               string     `json:"openId"`               // 用户在appId下的唯一标识
	PayMethod            string     `json:"payMethod"`            // 支付方式 6：微信主扫支付 7：支付宝主扫支付 24：银联二维码主扫支付
}

type WxJSAPIPaymentRsp struct {
	ReturnCode   string        `json:"returnCode"`   // 返回状态码
	ReturnMsg    string        `json:"returnMsg"`    // 返回信息
	OutTradeNo   string        `json:"outTradeNo"`   // 商户订单号
	Amount       int64         `json:"amount"`       // 支付金额
	NonceStr     string        `json:"nonceStr"`     // 随机字符串
	WxJsapiParam *WxJsapiParam `json:"wxJsapiParam"` // JS中调起微信公众号支付所需的参数
}

type WxJsapiParam struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	SignType  string `json:"signType"`
	Package   string `json:"package"`
	PaySign   string `json:"paySign"`
}

// 微信公众号/小程序支付接口
func (c *Client) WxJSAPIPayment(wjsapip *WxJSAPIPayment) (rsp *WxJSAPIPaymentRsp, err error) {
	rsp = new(WxJSAPIPaymentRsp)

	err = wjsapip.checkParms()
	if err != nil {
		return rsp, err
	}

	req := newSetting(wxH5Payment, c)
	err = req.doPostReq(wjsapip, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (wjsapip *WxJSAPIPayment) checkParms() error {
	if wjsapip.OutTradeNo == "" {
		return errors.New("商户订单号不能为空")
	}
	if wjsapip.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if wjsapip.ClientIp == "" {
		return errors.New("用户终端IP不能为空")
	}
	if wjsapip.PayMethod != "6" && wjsapip.PayMethod != "7" && wjsapip.PayMethod != "24" {
		return errors.New("请选择扫描类型")
	}
	if wjsapip.AppId == "" {
		return errors.New("AppId不能为空")
	}
	if wjsapip.OpenId == "" {
		return errors.New("OpenId不能为空")
	}
	if wjsapip.OrderInfo == nil {
		return errors.New("商品订单信息不能为空")
	}
	if wjsapip.PayAmount <= 0 {
		return errors.New("支付金额小于等于0")
	}
	if wjsapip.PayCurrency == "" {
		return errors.New("支付币种不能为空")
	}
	if wjsapip.ChannelType != "01" && wjsapip.ChannelType != "02" {
		return errors.New("渠道类型必须为空01或者02")
	}
	if wjsapip.NotifyUrl == "" {
		return errors.New("支付结果通知地址为空")
	}
	if wjsapip.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	if wjsapip.TransactionStartTime == "" {
		return errors.New("交易开始时间必填")
	}
	return nil
}
