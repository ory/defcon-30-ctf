package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testHandlers(t *testing.T) http.Handler {
	t.Helper()
	repo, err := NewRepo(&Config{
		DataSourceName: "sqlite3://:memory:",
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
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/results/"+r.District, bytes.NewBuffer(bs))
	require.NoError(t, err)

	handler(w, req)
	return w
}

func TestSubmitAndGet(t *testing.T) {
	t.Parallel()
	api := testHandlers(t).ServeHTTP

	assert.HTTPBodyContains(t, api, "GET", "/results", nil, "[]")

	results := []*result{{
		District:    "District 1",
		Democrats:   1,
		Republicans: 1,
		Invalid:     0,
	}, {
		District:    "District 2",
		Democrats:   2,
		Republicans: 0,
		Invalid:     1,
	}}

	for _, r := range results {
		res := submit(t, api, r)
		assert.Equal(t, http.StatusCreated, res.Code, "%s", res.Body.String())
	}

	expected, _ := json.Marshal(results)
	actual := assert.HTTPBody(api, "GET", "/results", nil)
	assert.JSONEq(t, string(expected), actual)
}

func TestSubmitUniqueDistrictIDs(t *testing.T) {
	t.Parallel()
	api := testHandlers(t).ServeHTTP

	res := submit(t, api, &result{District: "foo"})
	assert.Equal(t, http.StatusCreated, res.Code, "first creat should succeed: %s", res.Body.String())

	res = submit(t, api, &result{District: "foo"})
	assert.Equal(t, http.StatusBadRequest, res.Code, "second create should fail: %s", res.Body.String())
}
