package hduhelp

import "time"

type ValidateResponse struct {
	AccessToken string `json:"accessToken"`
	ClientID    string `json:"clientID"`
	ExpiredTime int64  `json:"expiredTime"`
	GrantType   string `json:"grantType"`
	IsValid     int64  `json:"isValid"`
	IsViald     int64  `json:"isViald"` //适配之前的版本
	School      string `json:"school"`
	StaffID     string `json:"staffId"`
	TokenType   int64  `json:"tokenType"`
	Uid         int64  `json:"uid"`
}

func (resp ValidateResponse) Valid() bool {
	return resp.IsValid == 1
}

func (resp ValidateResponse) GetExpireTime() time.Time {
	return time.Unix(resp.ExpiredTime, 0)
}
