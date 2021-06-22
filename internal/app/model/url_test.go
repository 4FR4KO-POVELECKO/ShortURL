package model_test

import (
	"ShortURL/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		payload *model.URL
		err     bool
	}{
		{
			name: "valid",
			payload: &model.URL{
				OriginURL: "google.com",
				ShortURL:  "JL_Cfwf951",
			},
			err: false,
		},
		{
			name: "empty",
			payload: &model.URL{
				OriginURL: "",
				ShortURL:  "",
			},
			err: true,
		},
		{
			name: "invalid origin url",
			payload: &model.URL{
				OriginURL: "googlecom",
				ShortURL:  "XYZabc123_",
			},
			err: true,
		},
		{
			name: "short url < 10",
			payload: &model.URL{
				OriginURL: "google.com",
				ShortURL:  "JLCfwf_95",
			},
			err: true,
		},
		{
			name: "short url > 10",
			payload: &model.URL{
				OriginURL: "google.com",
				ShortURL:  "XYZabc1234_",
			},
			err: true,
		},
		{
			name: "short url without numbers",
			payload: &model.URL{
				OriginURL: "google.com",
				ShortURL:  "XYZabcdef_",
			},
			err: true,
		},
		{
			name: "short url without upper letters",
			payload: &model.URL{
				OriginURL: "google.com",
				ShortURL:  "abcdef123_",
			},
			err: true,
		},
		{
			name: "short url without lower letters",
			payload: &model.URL{
				OriginURL: "google.com",
				ShortURL:  "XYZABC123_",
			},
			err: true,
		},
		{
			name: "short url without underscore",
			payload: &model.URL{
				OriginURL: "google.com",
				ShortURL:  "XYZabc1234",
			},
			err: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.payload.ValidateURL()
			eq := false

			if err != nil {
				eq = assert.Error(t, err)
			}

			assert.Equal(t, tc.err, eq)
		})
	}
}
