package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testHandlers(t *testing.T) http.Handler {
	t.Helper()
	repo, err := NewRepo(&Config{
		DataSourceName: ":memory:",
	})
	if err != nil {
		t.Fatal(err)
	}
	return NewHandler(repo)
}

func submit(t *testing.T, handler http.HandlerFunc, r *result) *httptest.ResponseRecorder {
	t.Helper()
	w := httptest.NewRecorder()
	bs, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/results", bytes.NewBuffer(bs))
	if err != nil {
		t.Fatal(err)
	}

	handler(w, req)
	return w
}

func TestSubmitAndGet(t *testing.T) {
	t.Parallel()
	api := testHandlers(t).ServeHTTP

	assert.HTTPBodyContains(t, api, "GET", "/results", nil, "[]")

	results := []*result{{
		DistrictID: 1,
		Votes: map[string]uint{
			"party a":    12,
			"party b":    34,
			"the answer": 42,
		},
	}, {
		DistrictID: 2,
		Votes: map[string]uint{
			"party a":    1,
			"party b":    1,
			"the answer": 0,
		},
	}}

	for _, r := range results {
		res := submit(t, api, r)
		assert.Equal(t, http.StatusCreated, res.Code)
	}

	expected, _ := json.Marshal(results)
	actual := assert.HTTPBody(api, "GET", "/results", nil)
	assert.JSONEq(t, string(expected), actual)
}

func TestSubmitUniqueDistrictIDs(t *testing.T) {
	t.Parallel()
	api := testHandlers(t).ServeHTTP

	res := submit(t, api, &result{DistrictID: 123})
	assert.Equal(t, http.StatusCreated, res.Code, "first create succeeds")

	res = submit(t, api, &result{DistrictID: 123})
	assert.Equal(t, http.StatusBadRequest, res.Code, "second create fails")
}
