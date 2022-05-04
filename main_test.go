package main

import (
	"fmt"
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

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected ststus code 200, got %v", err)
	}

	responseData, _ := io.ReadAll(resp.Body)
	if string(responseData) != mockUserResp {
		t.Fatalf("Expect \n%v message \n, got \n%v", mockUserResp, string(responseData))
	}

}
