package epaylinks

import (
	"errors"
)

type ProtocolBindCard struct {
	MchtOrderNo      string `json:"mchtOrderNo"`      // 请求流水号
	CustomerCode     string `json:"customerCode"`     // 商户号
	MemberId         string `json:"memberId"`         // 会员编号
	UserName         string `json:"userName"`         // 会员姓名
	PhoneNum         string `json:"phoneNum"`         // 手机号码
	BankCardNo       string `json:"bankCardNo"`       // 银行卡号
	BankCardType     string `json:"bankCardType"`     // 银行卡类型
	CertificatesType string `json:"certificatesType"` // 证件类型
	CertificatesNo   string `json:"certificatesNo"`   // 证件号码
	Expired          string `json:"expired"`          // 有效日期
	Cvn              string `json:"cvn"`              //  CVN 码
	NonceStr         string `json:"nonceStr"`         // 随机字符串
}

type ProtocolBindCardRsp struct {
	ReturnCode   string `json:"returnCode"`   // 返回状态码
	ReturnMsg    string `json:"returnMsg"`    // 返回信息
	SmsNo        string `json:"smsNo"`        // 绑卡流水号
	CustomerCode string `json:"customerCode"` // 商户号
	MemberId     string `json:"memberId"`     // 会员编号
	NonceStr     string `json:"nonceStr"`     // 随机字符串
}

// 协议支付绑卡预交易
func (c *Client) ProtocolBindCard(pbc *ProtocolBindCard) (rsp *ProtocolBindCardRsp, err error) {
	rsp = new(ProtocolBindCardRsp)

	err = pbc.checkParms()
	if err != nil {
		return rsp, err
	}

	req := newSetting(protocolBindCard, c)
	err = req.doPostReq(pbc, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (pbc *ProtocolBindCard) checkParms() error {
	if pbc.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if pbc.MemberId == "" {
		return errors.New("会员编号不能为空")
	}
	if pbc.UserName == "" {
		return errors.New("会员姓名不能为空")
	}
	if pbc.PhoneNum == "" {
		return errors.New("手机号码不能为空")
	}
	if pbc.BankCardNo == "" {
		return errors.New("银行卡号不能为空")
	}
	if pbc.BankCardType != "debit" && pbc.BankCardType != "credit" {
		return errors.New("银行卡类型不能为空")
	}
	if pbc.CertificatesType != "01" {
		return errors.New("证件类型不能为空")
	}
	if pbc.CertificatesNo == "" {
		return errors.New("证件号码不能为空")
	}
	if pbc.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	return nil
}
