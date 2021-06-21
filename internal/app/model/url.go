package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type URL struct {
	OriginURL string `json:"origin_url"`
	ShortURL  string `json:"short_url"`
}

func (u *URL) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.OriginURL, validation.Required, is.URL),
		validation.Field(
			&u.ShortURL,
			validation.Required,
			validation.Length(10, 10),
			validation.Match(regexp.MustCompile("[A-Z]{1,3}")),
			validation.Match(regexp.MustCompile("[a-z]{1,3}")),
			validation.Match(regexp.MustCompile("[0-9]{1,3}")),
			validation.Match(regexp.MustCompile("_")),
		),
	)
}
