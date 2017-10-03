package spgateway

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
)

// Config provides spgateway config.
type Config struct {
	HashKey    string
	HashIV     string
	MerchantID string
}

// Store provides spgateway config.
type Store struct {
	HashKey    string
	HashIV     string
	MerchantID string
}

// OrderCheckValue provides order config.
type OrderCheckValue struct {
	Amt             int
	MerchantOrderNo string
	MerchantID      string
	TimeStamp       string
	Version         string
}

// OrderCheckCode provides order config.
type OrderCheckCode struct {
	Amt             int
	MerchantOrderNo string
	MerchantID      string
	TradeNo         string
}

// Credit provides credit config.
type Credit struct {
	Date       string
	UseInfo    string
	CreditInst string
	CreditRed  string
	MerchantID string
}

// Invoice provides invoice check code config.
type Invoice struct {
	InvoiceTransNo  string
	MerchantOrderNo string
	RandomNum       string
	TotalAmt        int
	MerchantID      string
}

// New returns a new empty handler.
func New(config Config) *Store {
	return &Store{
		MerchantID: config.MerchantID,
		HashKey:    config.HashKey,
		HashIV:     config.HashIV,
	}
}

func (s *Store) hashSha256(str string) string {
	return strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256([]byte(str))))
}

// OrderCheckValue return spgateway check value for post data.
func (s *Store) OrderCheckValue(order OrderCheckValue) string {
	order.MerchantID = s.MerchantID
	v, _ := query.Values(order)

	querys := fmt.Sprintf("HashKey=%s&%s&HashIV=%s",
		s.HashKey,
		v.Encode(),
		s.HashIV,
	)

	return s.hashSha256(querys)
}

// OrderCheckCode return spgateway check value for post data.
func (s *Store) OrderCheckCode(order OrderCheckCode) string {
	order.MerchantID = s.MerchantID
	v, _ := query.Values(order)

	querys := fmt.Sprintf("HashIV=%s&%s&HashKey=%s",
		s.HashIV,
		v.Encode(),
		s.HashKey,
	)

	return s.hashSha256(querys)
}

// CreditCheckCode return spgateway check value for post data.
func (s *Store) CreditCheckCode(credit Credit) string {
	credit.MerchantID = s.MerchantID
	v, _ := query.Values(credit)

	querys := fmt.Sprintf("HashIV=%s&%s&HashKey=%s",
		s.HashIV,
		v.Encode(),
		s.HashKey,
	)

	return s.hashSha256(querys)
}

// InvoiceCheckCode return spgateway check value for post data.
func (s *Store) InvoiceCheckCode(invoice Invoice) string {
	invoice.MerchantID = s.MerchantID
	v, _ := query.Values(invoice)

	querys := fmt.Sprintf("HashIV=%s&%s&HashKey=%s",
		s.HashIV,
		v.Encode(),
		s.HashKey,
	)

	return s.hashSha256(querys)
}

// TradeSha return spgateway trade sha 256 encrypt.
func (s *Store) TradeSha(tradeInfo string) string {
	querys := fmt.Sprintf("HashKey=%s&%s&HashIV=%s",
		s.HashKey,
		tradeInfo,
		s.HashIV,
	)

	return s.hashSha256(querys)
}
