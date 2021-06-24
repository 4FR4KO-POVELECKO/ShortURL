package apiserver

import (
	"ShortURL/pkg/api"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/grpc"
)

func TestServer_HandleCreate(t *testing.T) {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewShortlinkClient(conn)
	s := newServer(c)

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]interface{}{
				"url": "google.com",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid url",
			payload: map[string]interface{}{
				"url": "googlecom",
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/create", b)
			s.ServeHTTP(rec, req)
			log.Println(rec.Code)
			//assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
