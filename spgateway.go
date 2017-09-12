package spgateway

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
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

// Order provides order config.
type Order struct {
	Amt             int
	MerchantOrderNo string
	TimeStamp       string
	Version         string
	TradeNo         string
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
func (s *Store) OrderCheckValue(order Order) string {
	querys := fmt.Sprintf("HashKey=%s&Amt=%s&MerchantID=%s&MerchantOrderNo=%s&TimeStamp=%s&Version=%s&HashIV=%s",
		s.HashKey,
		strconv.Itoa(order.Amt),
		s.MerchantID,
		order.MerchantOrderNo,
		order.TimeStamp,
		order.Version,
		s.HashIV,
	)

	return s.hashSha256(querys)
}

// OrderCheckCode return spgateway check value for post data.
func (s *Store) OrderCheckCode(order Order) string {
	querys := fmt.Sprintf("HashIV=%s&Amt=%s&MerchantID=%s&MerchantOrderNo=%s&TradeNo=%s&HashKey=%s",
		s.HashIV,
		strconv.Itoa(order.Amt),
		s.MerchantID,
		order.MerchantOrderNo,
		order.TradeNo,
		s.HashKey,
	)

	return s.hashSha256(querys)
}
