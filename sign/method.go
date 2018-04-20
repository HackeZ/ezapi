package sign

// Method of sign hashing algorithm
type Method string

const (
	METH_HMAC Method = "hmac"
	METH_MD5  Method = "md5"
)
