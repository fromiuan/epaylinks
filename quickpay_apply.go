package epaylinks

import (
	"errors"
)

type QuickpayApply struct {
	OutTradeNo           string     `json:"outTradeNo"`           // 商户订单号
	CustomerCode         string     `json:"customerCode"`         // 商户号
	ClientIp             string     `json:"clientIp"`             // 用户的IP
	OrderInfo            *OrderInfo `json:"orderInfo"`            // 商品订单信息
	PayAmount            int64      `json:"payAmount"`            // 支付金额
	PayCurrency          string     `json:"payCurrency"`          // 支付币种
	ChannelType          string     `json:"channelType"`          // 渠道类型
	AttachData           string     `json:"attachData"`           // 附加数据
	TransactionStartTime string     `json:"transactionStartTime"` // 交易开始时间
	NonceStr             string     `json:"nonceStr"`             // 随机字符串
	NeedSplit            bool       `json:"needSplit"`            // 是否分账
	RechargeMemCustCode  string     `json:"rechargeMemCustCode"`  // 充值会员客户编码
	AccountNo            string     `json:"accountNo"`            // 银行卡号码 String(32) 是 付款人银行卡号
	TelephoneNo          string     `json:"telephoneNo"`          // 持卡人银行预留手机号String(10) 是
	CertId               string     `json:"certId"`               // 身份证号 String(18) 是
	CertName             string     `json:"certName"`             // 身份证姓名 String(32) 是
	Cvn                  string     `json:"cvn"`                  // CVN 码 String(10) 否 信用卡必填
	ExpiredDate          string     `json:"expiredDate"`          // 信用卡过期时间 String(4) 否 信用卡必填 yymm，例如
}

type QuickpayApplyRsp struct {
	ReturnCode string `json:"returnCode"` // 返回状态码
	ReturnMsg  string `json:"returnMsg"`  // 返回信息
	OutTradeNo string `json:"outTradeNo"` // 商户订单号
	NonceStr   string `json:"nonceStr"`   // 随机字符串
	Token      string `json:"token"`      // 成功时必填，提交快捷支付时必填
}

// 申请快捷支付接口
func (c *Client) QuickpayApply(qa *QuickpayApply) (rsp *QuickpayApplyRsp, err error) {
	rsp = new(QuickpayApplyRsp)

	err = qa.checkParms()
	if err != nil {
		return rsp, err
	}

	req := newSetting(quickpayApply, c)
	err = req.doPostReq(qa, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (qa *QuickpayApply) checkParms() error {
	if qa.OutTradeNo == "" {
		return errors.New("商户订单号不能为空")
	}
	if qa.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if qa.ClientIp == "" {
		return errors.New("用户终端IP不能为空")
	}
	if qa.OrderInfo == nil {
		return errors.New("商品订单信息不能为空")
	}
	if qa.AccountNo == "" {
		return errors.New("银行卡号码不能为空")
	}
	if qa.TelephoneNo == "" {
		return errors.New("持卡人银行预留手机号不能为空")
	}
	if qa.CertId == "" {
		return errors.New("身份证号不能为空")
	}
	if qa.CertName == "" {
		return errors.New("身份证姓名不能为空")
	}
	if qa.PayAmount <= 0 {
		return errors.New("支付金额小于等于0")
	}
	if qa.PayCurrency == "" {
		return errors.New("支付币种不能为空")
	}
	if qa.ChannelType != "01" && qa.ChannelType != "02" {
		return errors.New("渠道类型必须为空01或者02")
	}
	if qa.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	if qa.TransactionStartTime == "" {
		return errors.New("交易开始时间必填")
	}
	return nil
}
