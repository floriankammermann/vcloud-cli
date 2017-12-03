package vcdapi

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestGetAllocatedIpsForNetworkHref(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("x-vcloud-authorization", "token")
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	testUrl := ts.URL
	GetAuthToken(testUrl, "user", "pass", "org")
	if len(vcdClient.VAToken) == 0 {
		t.Errorf("empty auth token")
	}
}
