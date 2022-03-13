package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var profileValidator ProfileValidator

func init() {
	profileValidator = ProfileValidator{}
}

func TestProfileValidator_ValidEmail(t *testing.T) {
	t.Parallel()
	err := profileValidator.ValidEmail("a@a.ru")
	assert.Nil(t, err)
	err = profileValidator.ValidEmail("test@test.com")
	assert.Nil(t, err)
	err = profileValidator.ValidEmail("testtest.com")
	assert.NotNil(t, err)
	err = profileValidator.ValidEmail("test@testcom")
	assert.NotNil(t, err)
	err = profileValidator.ValidEmail("testtestcom")
	assert.NotNil(t, err)
}

func TestPassword_ValidPassword(t *testing.T) {
	t.Parallel()
	pass := "test"
	err := profileValidator.ValidPassword(pass)
	assert.NotNil(t, err)
	pass = "testtest"
	err = profileValidator.ValidPassword(pass)
	assert.NotNil(t, err)
	pass = "12345678"
	err = profileValidator.ValidPassword(pass)
	assert.NotNil(t, err)
	pass = "test1test"
	err = profileValidator.ValidPassword(pass)
	assert.Nil(t, err)

}

func TestPhone_ValidPhone(t *testing.T) {
	t.Parallel()
	var p uint64 = 79062579330
	err := profileValidator.ValidPhone(p)
	assert.Nil(t, err)
	p = 9062579330
	err = profileValidator.ValidPhone(p)
	assert.NotNil(t, err)
}
