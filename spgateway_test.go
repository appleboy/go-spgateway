package spgateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderCheckValue(t *testing.T) {
	store := New(Config{
		MerchantID: "123456",
		HashKey:    "1A3S21DAS3D1AS65D1",
		HashIV:     "1AS56D1AS24D",
	})

	order := OrderCheckValue{
		Amt:             200,
		MerchantOrderNo: "20140901001",
		TimeStamp:       "1403243286",
		Version:         "1.1",
	}

	CheckValue := store.OrderCheckValue(order)
	expect := "841F57D750FB4B04B62DDC3ECDC26F1F4028410927DD28BD5B2E34791CC434D2"

	assert.Equal(t, expect, CheckValue)
}

func TestOrderCheckCode(t *testing.T) {
	store := New(Config{
		MerchantID: "1422967",
		HashKey:    "abcdefg",
		HashIV:     "1234567",
	})

	order := OrderCheckCode{
		Amt:             100,
		MerchantOrderNo: "840f022",
		TradeNo:         "14061313541640927",
	}

	CheckValue := store.OrderCheckCode(order)
	expect := "62C687AF6409E46E79769FAF54F54FE7E75AAE50BAF0767752A5C337670B8EDB"

	assert.Equal(t, expect, CheckValue)
}

func TestCreditCheckCode(t *testing.T) {
	store := New(Config{
		MerchantID: "ABC1422967",
		HashKey:    "abcdefg",
		HashIV:     "1234567",
	})

	credit := Credit{
		Date:       "2015-01-01 00:00:00",
		UseInfo:    "ON",
		CreditInst: "ON",
		CreditRed:  "ON",
	}

	CheckValue := store.CreditCheckCode(credit)
	expect := "77A1EF8F23C94CB63A60A7EDF99AC3E0F4688D96AF6D4B34370D306ABD33D0F6"

	assert.Equal(t, expect, CheckValue)
}

func TestInvoiceCheckCode(t *testing.T) {
	store := New(Config{
		MerchantID: "3622183",
		HashKey:    "abcdefg",
		HashIV:     "1234567",
	})

	invoice := Invoice{
		MerchantOrderNo: "201409170000001",
		InvoiceTransNo:  "14061313541640927",
		TotalAmt:        500,
		RandomNum:       "0142",
	}

	CheckValue := store.InvoiceCheckCode(invoice)
	expect := "C4156CA208897278C84D929DE48F4A2BCD1FF3ED4B97D09A14E2E2143E3EFD2E"

	assert.Equal(t, expect, CheckValue)
}

func TestTradeSha(t *testing.T) {
	store := New(Config{
		HashKey: "12345678901234567890123456789012",
		HashIV:  "1234567890123456",
	})

	tradeInfo := "ff91c8aa01379e4de621a44e5f11f72e4d25bdb1a18242db6cef9ef07d80b0165e476fd1d9acaa53170272c82d122961e1a0700a7427cfa1cf90db7f6d6593bbc93102a4d4b9b66d9974c13c31a7ab4bba1d4e0790f0cbbbd7ad64c6d3c8012a601ceaa808bff70f94a8efa5a4f984b9d41304ffd879612177c622f75f4214fa"

	tradeSha := store.TradeSha(tradeInfo)
	expect := "EA0A6CC37F40C1EA5692E7CBB8AE097653DF3E91365E6A9CD7E91312413C7BB8"

	assert.Equal(t, expect, tradeSha)
}
