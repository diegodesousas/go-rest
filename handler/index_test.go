package handler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)

	expected := `{"project":"Challenge Bravo Golang"}`
	assert.Equal(t, rr.Body.String(), expected)
}
