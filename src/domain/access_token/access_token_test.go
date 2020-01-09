package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// how to run with terminal
// > cd to_this_directory
// > go test
// > go test -cover

func TestAccessTokenConstants(t *testing.T) {
	// first way to write a test unit
	//if expirationTime != 24 { t.Error("expiration time should be 24 hours") }

	// second way this is the better way to write a test unit
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should NOT be expired")
}
