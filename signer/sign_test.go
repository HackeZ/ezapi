package signer

import (
	"testing"
)

const (
	PUB_AK = "c4339c0730175faba1284b7acb24c9db"
	PUB_AS = "45CAC98AF0F736A3AB2087CE26970674"
)

func TestSign_New(t *testing.T) {

	failCases := []struct {
		ak     string
		sk     string
		method Method
	}{
		{
			"",
			"",
			METH_MD5,
		},
		{
			PUB_AK,
			"",
			METH_HMAC,
		},
		{
			"",
			PUB_AS,
			METH_MD5,
		},
	}

	for i, c := range failCases {
		_, err := New(c.ak, c.sk, c.method)
		if err == nil {
			t.Fatalf("[%d] fail cases not return error\n", i)
		}
	}

	succCases := []struct {
		ak     string
		sk     string
		method Method
	}{
		{
			PUB_AK,
			PUB_AS,
			METH_MD5,
		},
		{
			PUB_AK,
			PUB_AS,
			METH_HMAC,
		},
	}

	for i, c := range succCases {
		_, err := New(c.ak, c.sk, c.method)
		if err != nil {
			t.Fatalf("[%d] success cases return error: %v", i, err)
		}
	}
}

func TestSign_Sum(t *testing.T) {

	sign, err := New(PUB_AK, PUB_AS, METH_MD5)
	if err != nil {
		t.Fatal("failed to new sign: ", err)
	}

	cases := []struct {
		params map[string]string
		sign   string
	}{
		{
			map[string]string{
				"product_id": "10096",
				"fields":     "product_id,name,name_en,is_onsale",
			},
			"4784FAD2943B0A13B8578594542AE00F",
		},
	}

	for i, c := range cases {
		s, err := sign.Sum("ezbuy.product.get", "1.0", c.params)
		if err != nil {
			t.Fatalf("[%d] failed to sum sign: %v", i, err)
		}

		if s != c.sign {
			t.Fatalf("[%d] sign not corrent, want: [%s] got: [%s]", i, c.sign, s)
		}
	}
}
