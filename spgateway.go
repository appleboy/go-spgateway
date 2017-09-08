package spgateway

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// Handler provides spgateway config.
type Handler struct {
	HashKey string
	HashIV  string
}

// New returns a new empty handler.
func New(HashKey, HashIV string) *Handler {
	return &Handler{
		HashKey: HashKey,
		HashIV:  HashIV,
	}
}

// CheckValue return spgateway check value for post data.
func (h *Handler) CheckValue(Amt, MerchantID, MerchantOrderNo, TimeStamp, Version string) string {
	querys := fmt.Sprintf("HashKey=%s&Amt=%s&MerchantID=%s&MerchantOrderNo=%s&TimeStamp=%s&Version=%s&HashIV=%s",
		h.HashKey,
		Amt,
		MerchantID,
		MerchantOrderNo,
		TimeStamp,
		Version,
		h.HashIV,
	)
	hash := sha256.Sum256([]byte(querys))

	return strings.ToUpper(fmt.Sprintf("%x", hash))
}
