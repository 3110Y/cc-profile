package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var password PasswordService

func init() {
	password = PasswordService{}
}

func TestPasswordService_Encode(t *testing.T) {
	t.Parallel()
	p := "test"
	encode, err := password.Encode(p)
	assert.Nil(t, err)
	assert.Equal(
		t,
		"ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e"+
			"473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c"+
			"304cc5fa9ad8e6f57f50028a8ff",
		*encode,
	)
}
