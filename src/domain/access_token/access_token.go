package access_token

import (
	"time"
)

const (
	// let's suppose expire time is 24 hours
	expirationTime = 24
)

// ClientId --> untuk bedain user agent, untuk keperluan khusus misal:
// kalau dari web kasih akses token expired 24 jam (Web frontend - Client-Id: 123)
// tapi kalau dari android expired 12 jam (Android - Client-Id: 234)
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
