package spgateway_test

import (
	"fmt"

	spgateway "github.com/appleboy/go-spgateway"
)

// CheckValue 組合及加密方法
//
// 排序欄位字串並用&符號串聯起來將回傳資料其中的五個欄位，分別是訂單金額(Amt)、商店代號(MerchantID)、
// 商店自訂單號(MerchantOrderNo)、時間戳記(TimeStamp)、程式串接版本(Version)，且參數需照英文字母 A~Z 排序
// 若第一字母相同比較第二字母，以此類推。將串聯後的字串前後加上商店專屬加密 HashKey 與商店專屬加密 HashIV。
// 將串聯後的字串用 SHA256 壓碼後轉大寫。
func ExampleStore_OrderCheckValue() {
	store := spgateway.New(spgateway.Config{
		MerchantID: "123456",
		HashKey:    "1A3S21DAS3D1AS65D1",
		HashIV:     "1AS56D1AS24D",
	})

	order := spgateway.OrderCheckValue{
		Amt:             200,
		MerchantOrderNo: "20140901001",
		TimeStamp:       "1403243286",
		Version:         "1.1",
	}

	fmt.Println(store.OrderCheckValue(order))

	// Output:
	// 841F57D750FB4B04B62DDC3ECDC26F1F4028410927DD28BD5B2E34791CC434D2
}

// CheckCode 產生規則
//
// 排序欄位字串並用&符號串聯起來將回傳資料其中的四個欄位，分別是 Amt(金額)、MerchantID(商店代號)、
// MerchantOrderNo(商店訂單編號)、TradeNo(智付通交易序號)，且參數需照英文字母 A~Z 排序，若第一
// 字母相同比較第二字母，以此類推。將串聯後的字串前後加上商店專屬加密 HashIV 值與商店專屬加密 HashKey 值。
// 將串聯後的字串用 SHA256 壓碼後轉大寫。
func ExampleStore_OrderCheckCode() {
	store := spgateway.New(spgateway.Config{
		MerchantID: "1422967",
		HashKey:    "abcdefg",
		HashIV:     "1234567",
	})

	order := spgateway.OrderCheckCode{
		Amt:             100,
		MerchantOrderNo: "840f022",
		TradeNo:         "14061313541640927",
	}

	fmt.Println(store.OrderCheckCode(order))

	// Output:
	// 62C687AF6409E46E79769FAF54F54FE7E75AAE50BAF0767752A5C337670B8EDB
}
