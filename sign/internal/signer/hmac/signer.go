package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"

	"github.com/hackez/ezapi/sign/internal/key"
)

// HMAC generate sign using hmac by md5
type HMAC struct {
	sign []byte
}

// Get return sign after hmac by md5 encoding
func (h *HMAC) Get(k key.Key, args string) (string, error) {
	mac := hmac.New(md5.New, k.SecretKey)
	_, err := mac.Write([]byte(args))
	if err != nil {
		return "", err
	}

	h.sign = mac.Sum(nil)
	return h.String(), nil
}

// Name of HMAC
func (h *HMAC) Name() string {
	return "hmac"
}

func (h HMAC) String() string {
	return hex.EncodeToString(h.sign)
}
