package workflow

import "time"

type IDCode struct {
	Code     CodeInfo `json:"code"`
	ExpireAt int64    `json:"expireAt"`
	Identity Identity `json:"identity"`
}

func (code IDCode) GetExpireTime() time.Time {
	return time.Unix(code.ExpireAt, 0)
}

type CodeInfo struct {
	Code128 string `json:"code128"`
	QR      string `json:"qr"` //二维码数据

}

type Identity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
