package md5

import (
	"bytes"
	cmd5 "crypto/md5"
	"encoding/hex"

	"github.com/hackez/ezapi/signer/internal/key"
)

// MD5 generate sign using md5
type MD5 struct {
	sign []byte
}

// Get return sign after md5 encoding
func (md *MD5) Get(k key.Key, args string) (sign string, err error) {
	hash := cmd5.New()

	bs := bytes.NewBuffer(k.SecretKey)
	_, err = bs.WriteString(args)
	if err != nil {
		return "", err
	}
	_, err = bs.Write(k.SecretKey)
	if err != nil {
		return "", err
	}

	_, err = hash.Write(bs.Bytes())
	if err != nil {
		return "", err
	}

	md.sign = hash.Sum(nil)
	return md.String(), nil
}

// Name of MD5 signer
func (md *MD5) Name() string {
	return "md5"
}

func (md MD5) String() string {
	return hex.EncodeToString(md.sign)
}
