package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP handler that we can pass to our secureHeaders
	// middleware, which writes a 200 status code and an "OK" response body.
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Pass the mock HTTP handler to our secureHeaders middleware. Because
	// secureHeaders *returns* a http.Handler we can call its ServeHTTP()
	// method, passing in the http.ResponseRecorder and dummy http.Request to
	// execute it.
	secureHeaders(next).ServeHTTP(rr, r)

	// Call the Result() method on the http.ResponseRecorder to get the results
	// of the test.
	rs := rr.Result()

	// Check that the middleware has correctly set the Content-Security-Policy
	// header on the response.
	expectedValue := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	if rs.Header.Get("Content-Security-Policy") != expectedValue {
		t.Errorf("got %q; want %q", rs.Header.Get("Content-Security-Policy"), expectedValue)
	}

	// Check that the middleware has correctly set the Referrer-Policy
	// header on the response.
	expectedValue = "origin-when-cross-origin"
	if rs.Header.Get("Referrer-Policy") != expectedValue {
		t.Errorf("got %q; want %q", rs.Header.Get("Referrer-Policy"), expectedValue)
	}

	// Check that the middleware has correctly set the X-Content-Type-Options
	// header on the response.
	expectedValue = "nosniff"
	if rs.Header.Get("X-Content-Type-Options") != expectedValue {
		t.Errorf("got %q; want %q", rs.Header.Get("X-Content-Type-Options"), expectedValue)
	}

	// Check that the middleware has correctly set the X-Frame-Options header
	// on the response.
	expectedValue = "deny"
	if rs.Header.Get("X-Frame-Options") != expectedValue {
		t.Errorf("got %q; want %q", rs.Header.Get("X-Frame-Options"), expectedValue)
	}

	// Check that the middleware has correctly set the X-XSS-Protection header
	// on the response
	expectedValue = "0"
	if rs.Header.Get("X-XSS-Protection") != expectedValue {
		t.Errorf("got %q; want %q", rs.Header.Get("X-XSS-Protection"), expectedValue)
	}

	// Check that the middleware has correctly called the next handler in line
	// and the response status code and body are as expected.
	if rs.StatusCode != 200 {
		t.Errorf("got %q; want %q", rs.StatusCode, 200)
	}

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)

	if string(body) != "OK" {
		t.Errorf("got %q; want %q", string(body), "OK")
	}
}
