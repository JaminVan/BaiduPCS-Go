package panhome

import (
	"github.com/Luis0001/Baidu-Login/bdcrypto"
	"github.com/Luis0001/BaiduPCS-Go/baidupcs/netdisksign"
	"github.com/Luis0001/BaiduPCS-Go/pcsutil/converter"
)

type (
	// SignRes 签名结果
	SignRes interface {
		Sign() string
		Timestamp() string
	}

	signRes struct {
		sign      string
		timestamp string
	}
)

func (sr *signRes) Sign() string {
	return sr.sign
}
func (sr *signRes) Timestamp() string {
	return sr.timestamp
}

func (ph *PanHome) Signature() (sign SignRes, err error) {
	err = ph.getSignInfo()
	if err != nil {
		return nil, err
	}

	o := netdisksign.Sign2(ph.sign3, ph.sign1)
	signed := bdcrypto.Base64Encode(o)
	return &signRes{
		sign:      converter.ToString(signed),
		timestamp: ph.timestamp,
	}, nil
}
