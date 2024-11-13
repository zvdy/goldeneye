package main_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/zvdy/goldeneye/pkg/dos"
	"github.com/zvdy/goldeneye/pkg/options"
)

func TestGoldenEye(t *testing.T) {
	// Create multiple test servers with different response codes
	servers := []struct {
		server *httptest.Server
		method string
		status int
	}{
		{
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})),
			method: "get",
			status: http.StatusOK,
		},
		{
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			})),
			method: "post",
			status: http.StatusNotFound,
		},
		{
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			})),
			method: "put",
			status: http.StatusInternalServerError,
		},
		{
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusForbidden)
			})),
			method: "delete",
			status: http.StatusForbidden,
		},
	}

	for _, srv := range servers {
		defer srv.server.Close()

		// Set up options for the test
		opts := &options.Options{
			URL:        srv.server.URL,
			UserAgents: "randomly generated",
			Workers:    10,
			Sockets:    5,
			Method:     srv.method,
			Debug:      true,
			NoSSLCheck: true,
		}

		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Initialize and start the DoS attack
		goldeneye := dos.NewGoldenEye(opts.URL, opts.Workers, opts.Sockets, opts.Method, opts.Debug, opts.NoSSLCheck, opts.UserAgents)
		go goldeneye.Fire(ctx)

		// Wait for the context to timeout
		<-ctx.Done()

		// Since the Fire method does not return a response, we need to manually check the URL
		resp, err := http.Get(srv.server.URL)
		if err != nil {
			t.Fatalf("Failed to make GET request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != srv.status {
			t.Errorf("Expected status code %d, got %d", srv.status, resp.StatusCode)
		}
	}
}
