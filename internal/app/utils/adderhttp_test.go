package shorten_test

import (
	shorten "ShortURL/internal/app/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddHTTP(t *testing.T) {
	url := "google.com"
	matched := shorten.AddHTTP(url)
	assert.Equal(t, matched, "https://google.com")
}
