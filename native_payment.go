package epaylinks

import (
	"errors"
)

type NativePayment struct {
	OutTradeNo           string     `json:"outTradeNo"`           // 商户订单号
	CustomerCode         string     `json:"customerCode"`         // 商户号
	SubCustomerCode      string     `json:"subCustomerCode"`      // 子商户号
	TerminalCode         string     `json:"terminalCode"`         // 终端代码
	PayMethod            string     `json:"payMethod"`            // 支付方式 6：微信主扫支付 7：支付宝主扫支付 24：银联二维码主扫支付
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
}

type NativePaymentRsp struct {
	ReturnCode string `json:"returnCode"` // 返回状态码
	ReturnMsg  string `json:"returnMsg"`  // 返回信息
	CodeUrl    string `json:"codeUrl"`    // 扫码 URL
	MwebURL    string `json:"mwebURL"`    // 重定向URL
	OutTradeNo string `json:"outTradeNo"` // 商户订单号
	Amount     int64  `json:"amount"`     // 支付金额
	NonceStr   string `json:"nonceStr"`   // 随机字符串
}

// 主扫支付接口
func (c *Client) NativePayment(np *NativePayment) (rsp *NativePaymentRsp, err error) {
	rsp = new(NativePaymentRsp)

	err = np.checkParms()
	if err != nil {
		return rsp, err
	}

	err = c.doPostReq(nativePayment, np, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (np *NativePayment) checkParms() error {
	if np.OutTradeNo == "" {
		return errors.New("商户订单号不能为空")
	}
	if np.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if np.ClientIp == "" {
		return errors.New("用户终端IP不能为空")
	}
	if np.PayMethod != "6" && np.PayMethod != "7" && np.PayMethod != "24" {
		return errors.New("请选择扫描类型")
	}
	if np.OrderInfo == nil {
		return errors.New("商品订单信息不能为空")
	}
	if np.PayAmount <= 0 {
		return errors.New("支付金额小于等于0")
	}
	if np.PayCurrency == "" {
		return errors.New("支付币种不能为空")
	}
	if np.ChannelType != "01" && np.ChannelType != "02" {
		return errors.New("渠道类型必须为空01或者02")
	}
	if np.NotifyUrl == "" {
		return errors.New("支付结果通知地址为空")
	}
	if np.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	if np.TransactionStartTime == "" {
		return errors.New("交易开始时间必填")
	}
	return nil
}
