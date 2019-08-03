package epaylinks

import (
	"errors"
)

type ProtocolUnBindCard struct {
	CustomerCode string `json:"customerCode"` // 商户号
	MemberId     string `json:"memberId"`     // 会员编号
	Protocol     string `json:"protocol"`     // 协议号
	NonceStr     string `json:"nonceStr"`     // 随机字符串
}

type ProtocolUnBindCardRsp struct {
	ReturnCode   string `json:"returnCode"`   // 返回状态码
	ReturnMsg    string `json:"returnMsg"`    // 返回信息
	CustomerCode string `json:"customerCode"` // 商户号
	MemberId     string `json:"memberId"`     // 会员编号 是
	NonceStr     string `json:"nonceStr"`     // 随机字符串
}

// 协议支付解绑
func (c *Client) ProtocolUnBindCard(pubc *ProtocolUnBindCard) (rsp *ProtocolUnBindCardRsp, err error) {
	rsp = new(ProtocolUnBindCardRsp)

	err = pubc.checkParms()
	if err != nil {
		return rsp, err
	}

	err = c.doPostReq(protocolunBindCard, pubc, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (pubc *ProtocolUnBindCard) checkParms() error {
	if pubc.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if pubc.Protocol == "" {
		return errors.New("协议号不能为空")
	}
	if pubc.MemberId == "" {
		return errors.New("会员编号不能为空")
	}
	if pubc.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	return nil
}
