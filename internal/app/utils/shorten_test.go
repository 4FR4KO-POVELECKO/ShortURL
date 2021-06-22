package shorten_test

import (
	"ShortURL/internal/app/model"
	shorten "ShortURL/internal/app/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Shorten(t *testing.T) {
	short := shorten.Shorten()

	url := model.TestURL(t)
	url.ShortURL = short

	err := url.ValidateURL()
	assert.NoError(t, err)
}
