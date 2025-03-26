package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerCorrectRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
	cafes := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, cafes, 2)
}

func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=london&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow&count=10", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	cafes := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, cafes, 4)
}