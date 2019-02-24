package tests

import (
	"log"
	"testing"

	freya "github.com/anargu/freya/freyacon/go"
	a "github.com/stretchr/testify/assert"
)

func TestNewFreya(t *testing.T) {
	assert := a.New(t)
	log.Println(t.Name())
	f, err := freya.NewFreya(&freya.FreyaConfig{})
	assert.NotNil(f, "Freya not support empty config")
	assert.Nil(err, "Error creating freya empty config")
}
