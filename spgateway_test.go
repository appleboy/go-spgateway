package spgateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValue(t *testing.T) {
	h := New("1A3S21DAS3D1AS65D1", "1AS56D1AS24D")
	CheckValue := h.CheckValue("200", "123456", "20140901001", "1403243286", "1.1")

	assert.Equal(t, "841F57D750FB4B04B62DDC3ECDC26F1F4028410927DD28BD5B2E34791CC434D2", CheckValue)
}
