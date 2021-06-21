package apiserver

import (
	"github.com/asirkov/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", nil)

	s := newServer(teststore.New())

	s.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}
