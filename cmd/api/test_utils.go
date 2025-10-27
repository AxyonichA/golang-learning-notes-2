package main

import (
	"net/http"
	"net/http/httptest"
	"social/internal/auth"
	"social/internal/store"
	"social/internal/store/cache"
	"testing"

	"go.uber.org/zap"
)

func newTestApplication(t *testing.T) *application {
	t.Helper()

	logger := zap.NewNop().Sugar()
	mockStore := store.NewMockStore()
	mockCacheStore := cache.NewMockStore()
	testAuth := &auth.TestAuthenticator{}

	return &application{
		store:         mockStore,
		logger:        logger,
		cacheStorage:  mockCacheStore,
		authenticator: testAuth,
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}
