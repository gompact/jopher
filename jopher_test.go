package jopher

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSuccess(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		Success(w, M{"Hello": "World"})
	}

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body := make(map[string]string)
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Log(err.Error())
	}
	defer resp.Body.Close()

	// parse body as json
	if body["Hello"] != "World" {
		t.Errorf("failed to get correct values from response body")
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected response status code to be 200, found %d", resp.StatusCode)
		return
	}

	if !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		t.Errorf("Response header Content-Type should be application/json")
		return
	}
}

func TestCreated(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		Created(w, M{"Hello": "World"})
	}

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body := make(map[string]string)
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Log(err.Error())
	}
	defer resp.Body.Close()

	// parse body as json
	if body["Hello"] != "World" {
		t.Errorf("failed to get correct values from response body")
		return
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected response status code to be 201, found %d", resp.StatusCode)
		return
	}

	if !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		t.Errorf("Response header Content-Type should be application/json")
		return
	}
}