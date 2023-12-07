package userdata

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCompute2FACode(t *testing.T) {
	us := NewUserService(nil, logrus.New())
	testCases := []struct {
		description string
		sub1        string
		sub2        string
	}{
		{"items01", "J8SWY3DPK9XXE3DE", "JBSWY3DPK6XXE3DE"},
		{"items02", "JBSWY31PK5XXEooDE------", "JBSWY3iPK5XXE00DE"},
		{"empty values", " ", "ndks"},
		{"wrong base32", " ", "ooooooooooooooooooooooooooooooooooooooooooooo"},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, us.Compute2FACode(tc.sub1), us.Compute2FACode(tc.sub2), tc.description)
		})
	}

}
