package epaylinks

import (
	"errors"
)

type H5WxPayment struct {
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
	SceneInfo            *SceneInfo `json:"SceneInfo"`            // 场景信息
}

type SceneInfo struct {
	Type    string `json:"type"`    // 类型
	WapURL  string `json:"wapURL"`  // WAP网站 URL
	WapName string `json:"wapName"` // WAP网站名称
}

type H5WxPaymentRsp struct {
	ReturnCode string `json:"returnCode"` // 返回状态码
	ReturnMsg  string `json:"returnMsg"`  // 返回信息
	MwebURL    string `json:"mwebURL"`    // 重定向URL
	OutTradeNo string `json:"outTradeNo"` // 商户订单号
	Amount     int64  `json:"amount"`     // 支付金额
	NonceStr   string `json:"nonceStr"`   // 随机字符串
}

// 微信H5支付接口
func (c *Client) H5WxPayment(hwp *H5WxPayment) (rsp *H5WxPaymentRsp, err error) {
	rsp = new(H5WxPaymentRsp)

	err = hwp.checkParms()
	if err != nil {
		return rsp, err
	}

	req := newSetting(wxH5Payment, c)
	err = req.doPostReq(hwp, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (hwp *H5WxPayment) checkParms() error {
	if hwp.OutTradeNo == "" {
		return errors.New("商户订单号不能为空")
	}
	if hwp.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if hwp.ClientIp == "" {
		return errors.New("用户终端IP不能为空")
	}
	if hwp.OrderInfo == nil {
		return errors.New("商品订单信息不能为空")
	}
	if hwp.PayAmount <= 0 {
		return errors.New("支付金额小于等于0")
	}
	if hwp.PayCurrency == "" {
		return errors.New("支付币种不能为空")
	}
	if hwp.ChannelType != "01" && hwp.ChannelType != "02" {
		return errors.New("渠道类型必须为空01或者02")
	}
	if hwp.NotifyUrl == "" {
		return errors.New("支付结果通知地址为空")
	}
	if hwp.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	if hwp.TransactionStartTime == "" {
		return errors.New("交易开始时间必填")
	}
	return nil
}
