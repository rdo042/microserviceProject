package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewBalance(t *testing.T) {
	balance := NewBalabce("11111111", "222222", 10, 20)

	assert.Nil(t, err)
	assert.NotNil(t, balance)

}
