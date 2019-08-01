package epaylinks

import (
	"errors"
)

type MicroPayment struct {
	OutTradeNo           string     `json:"outTradeNo"`           // 商户订单号
	CustomerCode         string     `json:"customerCode"`         // 商户号
	SubCustomerCode      string     `json:"subCustomerCode"`      // 子商户号
	TerminalCode         string     `json:"terminalCode"`         // 终端代码
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
	Scene                string     `json:"scene"`                // 场景信息 bar_code：条码支付 wave_code：声波支付
	AuthCode             string     `json:"authCode"`             // 付款授权码
}

type MicroPaymentRsp struct {
	ReturnCode        string `json:"returnCode"`        // 返回状态码
	ReturnMsg         string `json:"returnMsg"`         // 返回信息
	OutTradeNo        string `json:"outTradeNo"`        // 商户订单号
	Amount            int64  `json:"amount"`            // 支付金额
	NonceStr          string `json:"nonceStr"`          // 随机字符串
	TransactionNo     string `json:"transactionNo"`     // 易票联订单号 String 否 支付成功时必需
	ChannelOrder      string `json:"channelOrder"`      // 上游订单号 String 否 支付成功时必需
	ProcedureFee      int64  `json:"procedureFee"`      // 手续费 Long 否
	PayState          string `json:"payState"`          // 支付结果 String(2) 是
	PayTime           string `json:"payTime"`           // 支付完成时间 String(14) 否
	SettCycle         string `json:"settCycle"`         // 该支付交易所属的清算周期
	SettCycleInterval int    `json:"settCycleInterval"` //  清算周期长度
}

// 扫码（被扫）支付接口
func (c *Client) MicroPayment(mp *MicroPayment) (rsp *MicroPaymentRsp, err error) {
	rsp = new(MicroPaymentRsp)

	err = mp.checkParms()
	if err != nil {
		return rsp, err
	}

	req := newSetting(nativePayment, c)
	err = req.doPostReq(mp, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}
func (mp *MicroPayment) checkParms() error {
	if mp.OutTradeNo == "" {
		return errors.New("商户订单号不能为空")
	}
	if mp.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if mp.ClientIp == "" {
		return errors.New("用户终端IP不能为空")
	}
	if mp.AuthCode == "" {
		return errors.New("请填写付款授权码")
	}
	if mp.OrderInfo == nil {
		return errors.New("商品订单信息不能为空")
	}
	if mp.PayAmount <= 0 {
		return errors.New("支付金额小于等于0")
	}
	if mp.PayCurrency == "" {
		return errors.New("支付币种不能为空")
	}
	if mp.ChannelType != "01" && mp.ChannelType != "02" {
		return errors.New("渠道类型必须为空01或者02")
	}
	if mp.NotifyUrl == "" {
		return errors.New("支付结果通知地址为空")
	}
	if mp.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	if mp.TransactionStartTime == "" {
		return errors.New("交易开始时间必填")
	}
	return nil
}
