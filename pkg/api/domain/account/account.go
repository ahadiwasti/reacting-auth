package account

import (
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/scrypt"
)

const pwHashBytes = 64

// login type
var (
	LoginStandard = 1
	LoginOAuth    = 2
	LoginLdap     = 3
)

// login oauth type
const (
	OAuthDingTalk = iota
	//todo : ...
	OAuthWechat
	OAuthQQ
	OAuthFacebook
	OAuthGoogle
)

// HashPassword : password hashing
func HashPassword(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, pwHashBytes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h), nil
}

// MakeSalt : make password more complicated
func MakeSalt() (salt string, err error) {
	buf := make([]byte, pwHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", buf), nil
}
