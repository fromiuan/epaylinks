package epaylinks

import (
	"errors"
)

type ProtocolPayPre struct {
	OutTradeNo           string     `json:"outTradeNo"`           // 商户订单号
	CustomerCode         string     `json:"customerCode"`         // 商户号
	MemberId             string     `json:"memberId"`             // 会员编号
	Protocol             string     `json:"protocol"`             // 协议号
	SmsNo                string     `json:"smsNo"`                // 绑卡流水号
	SmsCode              string     `json:"smsCode"`              // 短信验证码
	OrderInfo            *OrderInfo `json:"orderInfo"`            // 商品订单信息
	PayAmount            int64      `json:"payAmount"`            // 支付金额
	PayCurrency          string     `json:"payCurrency"`          // 支付币种
	NotifyUrl            string     `json:"notifyUrl"`            // 支付结果通知地址
	AttachData           string     `json:"attachData"`           // 附加数据
	TransactionStartTime string     `json:"transactionStartTime"` // 交易开始时间
	NonceStr             string     `json:"nonceStr"`             // 随机字符串
	NeedSplit            bool       `json:"needSplit"`            // 是否分账
}

type ProtocolPayPreRsp struct {
	ReturnCode   string `json:"returnCode"`   // 返回状态码
	ReturnMsg    string `json:"returnMsg"`    // 返回信息
	OutTradeNo   string `json:"outTradeNo"`   // 商户订单号
	CustomerCode string `json:"customerCode"` // 商户号
	MemberId     string `json:"memberId"`     // 会员编号 是
	token        string `json:"token"`        // 唯一代表该订单
	NonceStr     string `json:"nonceStr"`     // 随机字符串
}

// 申请快捷支付接口
func (c *Client) ProtocolPayPre(upm *ProtocolPayPre) (rsp *ProtocolPayPreRsp, err error) {
	rsp = new(ProtocolPayPreRsp)

	err = upm.checkParms()
	if err != nil {
		return rsp, err
	}

	req := newSetting(protocolPayPre, c)
	err = req.doPostReq(upm, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (upm *ProtocolPayPre) checkParms() error {
	if upm.OutTradeNo == "" {
		return errors.New("商户订单号不能为空")
	}
	if upm.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if upm.OrderInfo == nil {
		return errors.New("商品订单信息不能为空")
	}
	if upm.Protocol == "" {
		return errors.New("协议号不能为空")
	}
	if upm.MemberId == "" {
		return errors.New("会员编号不能为空")
	}
	if upm.PayAmount <= 0 {
		return errors.New("支付金额小于等于0")
	}
	if upm.PayCurrency == "" {
		return errors.New("支付币种不能为空")
	}
	if upm.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	if upm.TransactionStartTime == "" {
		return errors.New("交易开始时间必填")
	}
	return nil
}
