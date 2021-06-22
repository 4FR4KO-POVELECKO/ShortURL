package shorten_test

import (
	"ShortURL/internal/app/model"
	shorten "ShortURL/internal/app/utils"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Shorten(t *testing.T) {
	short := shorten.Shorten()
	url := model.URL{ShortURL: short}

	err := url.ValidateShortURL()
	assert.NoError(t, err)
	assert.Equal(t, reflect.TypeOf("string"), reflect.TypeOf(short))
}
