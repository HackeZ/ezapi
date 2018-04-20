package signer

import (
	"github.com/hackez/ezapi/sign/internal/key"
	"github.com/hackez/ezapi/sign/internal/signer/hmac"
	"github.com/hackez/ezapi/sign/internal/signer/md5"
)

// Signer generate sign in different way
type Signer interface {
	Name() string
	Get(key key.Key, args string) (string, error)
}

var (
	// all implementation here
	_ Signer = &hmac.HMAC{}
	_ Signer = &md5.MD5{}
)
