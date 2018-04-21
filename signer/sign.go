package signer

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/hackez/ezapi/signer/internal/key"
	"github.com/hackez/ezapi/signer/internal/signer"
	"github.com/hackez/ezapi/signer/internal/signer/hmac"
	"github.com/hackez/ezapi/signer/internal/signer/md5"
)

const (
	FMT_TIMESTAMP = "2006-01-02 15:04:05"
)

// Sign key in each request to ezbuy open api
type Sign struct {
	key    key.Key
	signer signer.Signer
}

// Method of sign hashing algorithm
type Method string

const (
	METH_HMAC Method = "hmac"
	METH_MD5  Method = "md5"
)

// New return available signer to sum sign
func New(appKey, secretKey string, method Method) (*Sign, error) {
	if appKey == "" || secretKey == "" {
		return nil, fmt.Errorf("invalid appKey or secretKey")
	}

	sign := &Sign{
		key: key.Key{
			AppKey:    []byte(appKey),
			SecretKey: []byte(secretKey),
		},
	}

	switch method {
	case METH_HMAC:
		sign.signer = &hmac.HMAC{}
	case METH_MD5:
		sign.signer = &md5.MD5{}
	}

	return nil, fmt.Errorf("unknown method of sign")
}

// Sum available sign by api method/version and bussiness params
func (s *Sign) Sum(method, version string, params map[string]string) (string, error) {
	args := make(map[string]string)

	// public args
	args["method"] = method
	args["app_key"] = string(s.key.AppKey)
	args["sign_method"] = s.signer.Name()
	args["timestamp"] = time.Now().Format(FMT_TIMESTAMP)
	args["version"] = version

	// bussiness args if not null
	if params != nil {
		for k := range params {
			args[k] = params[k]
		}
	}

	token := argsFilter(args)
	token = url.QueryEscape(token)

	sign, err := s.signer.Get(s.key, token)
	if err != nil {
		return "", err
	}

	sign = strings.ToUpper(sign)
	return sign, nil
}

func argsFilter(args map[string]string) string {
	list := make([]string, 0, len(args))
	for k, v := range args {
		if k == "sign" { // ignore args of sign
			continue
		}
		list = append(list, fmt.Sprintf("%s=%s", k, v))
	}

	// sort by lexicographical order
	sort.Strings(list)

	return strings.Join(list, "")
}
