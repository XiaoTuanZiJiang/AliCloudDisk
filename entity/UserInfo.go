package entity

import (
	"bytes"
	"encoding/json"
)

type UserInfo struct {
	DomainId       string `json:"domain_id"`
	UserId         string `json:"user_id"`
	Avatar         string `json:"avatar"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
	LastLoginTime  int64  `json:"last_login_time"`
	Email          string `json:"email"`
	NickName       string `json:"nick_name"`
	Phone          string `json:"phone"`
	Status         string `json:"status"`
	DefaultDriveId string `json:"default_drive_id"`
}

func (u *UserInfo) String() string {
	bs, _ := json.Marshal(u)
	var out bytes.Buffer
	err := json.Indent(&out, bs, "", "\t")
	if err != nil {
		return ""
	}
	return out.String()
}
