package model

import "testing"

func TestURL(t *testing.T) *URL {
	t.Helper()

	return &URL{
		OriginURL: "www.google.com",
		ShortURL:  "XYZabc123_",
	}
}
