package main

import (
	"bytes"
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

func TestNewRecipeHandler(t *testing.T) {

	r := SetupRouter()
	r.POST("/recipes", NewRecipeHandler)
	recipes := Recipe{
		Name: "New York Pizza",
	}
	jsonValue, _ := json.Marshal(recipes)
	req, _ := http.NewRequest("POST", "/recipes", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateRecipeHandler(t *testing.T) {
	r := SetupRouter()
	r.PUT("/recipes/:id", UpdateRecipeHandler)
	recipe := Recipe{
		ID:   "c0283p3d0cvuglq85lpg",
		Name: "Gnocchi",
		Ingredients: []string{
			"5 large Idaho potatoes",
			"2 egges",
			"3/4 cup grated Parmesan",
			"3 1/2 cup all-purpose flour",
		},
	}
	jsonValue, _ := json.Marshal(recipe)
	reqFound, _ := http.NewRequest("PUT",
		"/recipes/"+recipe.ID, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)
	reqNotFound, _ := http.NewRequest("PUT", "/recipes/1",
		bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
