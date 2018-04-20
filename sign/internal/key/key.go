package key

import (
	"errors"
	"fmt"
)

// Key include AppKey and SecretKey. You
// can apply to ezbuy technicians to get them.
type Key struct {
	AppKey    []byte
	SecretKey []byte
}

// New available key if validate
func New(ak, sk string) (*Key, error) {
	if ak == "" || sk == "" {
		return nil, errors.New("invalid appKey or secretKey")
	}

	return &Key{
		AppKey:    []byte(ak),
		SecretKey: []byte(sk),
	}, nil
}

func (k Key) String() string {
	return fmt.Sprintf("AppKey: %s , SecretKey: %s", string(k.AppKey), string(k.SecretKey))
}
