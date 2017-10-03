package spgateway

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
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

// CreateMpgAesEncrypt encrypt trade info data
func (s *Store) CreateMpgAesEncrypt(tradeInfo interface{}) (string, error) {
	v, _ := query.Values(tradeInfo)
	data := []byte(v.Encode())
	key := []byte(s.HashKey)
	iv := []byte(s.HashIV)

	ciphertext, err := Encrypt(key, data, iv)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", ciphertext), nil
}

// Encrypt string
func Encrypt(key, plaintext, iv []byte) ([]byte, error) {
	plaintext = PKCS5Padding(plaintext, 32)
	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	return ciphertext, nil
}

// PKCS5Padding is described in RFC 5652.
func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// PKCS5UnPadding is described in RFC 5652.
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
