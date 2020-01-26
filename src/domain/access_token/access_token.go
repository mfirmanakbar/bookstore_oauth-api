package access_token

import (
	"github.com/mfirmanakbar/bookstore_oauth-api/src/utils/errors"
	"strings"
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

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
