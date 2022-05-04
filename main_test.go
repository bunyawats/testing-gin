package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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

func SetupRouter() *gin.Engine {

	router := gin.Default()
	return router
}

func TestListRecipesHandler(t *testing.T) {

	r := SetupRouter()
	r.GET("/recipes", ListRecipesHandler)
	req, _ := http.NewRequest("GET", "/recipes", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var recipes []Recipe
	json.Unmarshal([]byte(w.Body.String()), &recipes)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 492, len(recipes))
}
