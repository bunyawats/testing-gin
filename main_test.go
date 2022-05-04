package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {

	mockUserResp := `{"message":"hello world"}`
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	responseData, _ := io.ReadAll(resp.Body)
	assert.Equal(t, mockUserResp, string(responseData))
}
