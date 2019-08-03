package epaylinks

import (
	"errors"
)

type UnionPayMent struct {
	OutTradeNo           string     `json:"outTradeNo"`           // 商户订单号
	CustomerCode         string     `json:"customerCode"`         // 商户号
	OrderInfo            *OrderInfo `json:"orderInfo"`            // 商品订单信息
	PayAmount            int64      `json:"payAmount"`            // 支付金额
	PayCurrency          string     `json:"payCurrency"`          // 支付币种
	ChannelType          string     `json:"channelType"`          // 渠道类型
	NoCreditCards        bool       `json:"noCreditCards"`        // 是否禁止信用卡支付
	NotifyUrl            string     `json:"notifyUrl"`            // 支付结果通知地址
	AttachData           string     `json:"attachData"`           // 附加数据
	TransactionStartTime string     `json:"transactionStartTime"` // 交易开始时间
	NonceStr             string     `json:"nonceStr"`             // 随机字符串
	NeedSplit            bool       `json:"needSplit"`            // 是否分账
	FrontUrl             string     `json:"frontUrl"`             // 前台通知地址
	FrontFailUrl         string     `json:"frontFailUrl"`         // 交易失败前台跳转地址
	IssInsCode           string     `json:"issInsCode"`           // 银行简码 String(10) 是 发卡机构编码（跳转到特定的网银界面），见 4.7.4个人网银支付支持银行简码
	TradeType            string     `json:"tradeType"`            // 交易类型 String 否 union_online：银联在线 personal_bank: 网 银 支付 b2b：企业网银 不传默认为网银支付
}

type UnionPayMentRsp struct {
	ReturnCode string `json:"returnCode"` // 返回状态码
	ReturnMsg  string `json:"returnMsg"`  // 返回信息
	OutTradeNo string `json:"outTradeNo"` // 商户订单号
	NonceStr   string `json:"nonceStr"`   // 随机字符串
	Amount     int64  `json:"amount"`     // 成功时必填，提交快捷支付时必填
	RespMsg    string `json:"respMsg"`    // 应答信息
}

// 网银支付接口
func (c *Client) UnionPayMent(upm *UnionPayMent) (rsp *UnionPayMentRsp, err error) {
	rsp = new(UnionPayMentRsp)

	err = upm.checkParms()
	if err != nil {
		return rsp, err
	}

	err = c.doPostReq(unionPayMent, upm, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (upm *UnionPayMent) checkParms() error {
	if upm.OutTradeNo == "" {
		return errors.New("商户订单号不能为空")
	}
	if upm.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if upm.OrderInfo == nil {
		return errors.New("商品订单信息不能为空")
	}
	if upm.IssInsCode == "" {
		return errors.New("银行简码不能为空")
	}
	if upm.PayAmount <= 0 {
		return errors.New("支付金额小于等于0")
	}
	if upm.PayCurrency == "" {
		return errors.New("支付币种不能为空")
	}
	if upm.ChannelType != "01" && upm.ChannelType != "02" {
		return errors.New("渠道类型必须为空01或者02")
	}
	if upm.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	if upm.TransactionStartTime == "" {
		return errors.New("交易开始时间必填")
	}
	return nil
}
