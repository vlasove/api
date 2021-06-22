package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vlasove/api/internal/app/model"
	"github.com/vlasove/api/internal/app/store/teststore"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "example@user.org",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "email",
				"password": "password",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    "email@test.com",
				"password": "two",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			_ = json.NewEncoder(b).Encode(tt.payload)

			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tt.expectedCode)
		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {

	u := model.TestUser(t)
	store := teststore.New()
	s := newServer(store)
	err := store.User().Create(u)
	if err != nil {
		t.Fatalf("can not create test user")
	}
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusAccepted,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			_ = json.NewEncoder(b).Encode(tt.payload)

			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tt.expectedCode)
		})
	}
}
