package spgateway

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

// Handler provides spgateway config.
type Handler struct {
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
}

// New returns a new empty handler.
func New(MerchantID, HashKey, HashIV string) *Handler {
	return &Handler{
		MerchantID: MerchantID,
		HashKey:    HashKey,
		HashIV:     HashIV,
	}
}

// CheckValue return spgateway check value for post data.
func (h *Handler) CheckValue(order Order) string {
	querys := fmt.Sprintf("HashKey=%s&Amt=%s&MerchantID=%s&MerchantOrderNo=%s&TimeStamp=%s&Version=%s&HashIV=%s",
		h.HashKey,
		strconv.Itoa(order.Amt),
		h.MerchantID,
		order.MerchantOrderNo,
		order.TimeStamp,
		order.Version,
		h.HashIV,
	)
	hash := sha256.Sum256([]byte(querys))

	return strings.ToUpper(fmt.Sprintf("%x", hash))
}
