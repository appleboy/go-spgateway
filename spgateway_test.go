package spgateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValue(t *testing.T) {
	store := New(Config{
		MerchantID: "123456",
		HashKey:    "1A3S21DAS3D1AS65D1",
		HashIV:     "1AS56D1AS24D",
	})

	order := Order{
		Amt:             200,
		MerchantOrderNo: "20140901001",
		TimeStamp:       "1403243286",
		Version:         "1.1",
	}

	CheckValue := store.OrderCheckValue(order)
	expect := "841F57D750FB4B04B62DDC3ECDC26F1F4028410927DD28BD5B2E34791CC434D2"

	assert.Equal(t, expect, CheckValue)
}

func TestCheckCode(t *testing.T) {
	store := New(Config{
		MerchantID: "1422967",
		HashKey:    "abcdefg",
		HashIV:     "1234567",
	})

	order := Order{
		Amt:             100,
		MerchantOrderNo: "840f022",
		TradeNo:         "14061313541640927",
	}

	CheckValue := store.OrderCheckCode(order)
	expect := "62C687AF6409E46E79769FAF54F54FE7E75AAE50BAF0767752A5C337670B8EDB"

	assert.Equal(t, expect, CheckValue)
}
