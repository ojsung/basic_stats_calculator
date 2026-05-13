package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestFormHandler(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		wantStatus  int
		wantContain string
	}{
		{
			name:        "GET returns form",
			method:      http.MethodGet,
			wantStatus:  http.StatusOK,
			wantContain: "<form",
		},
		{
			name:       "POST returns 405",
			method:     http.MethodPost,
			wantStatus: http.StatusMethodNotAllowed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", nil)
			w := httptest.NewRecorder()
			formHandler(w, req)
			if w.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.wantStatus)
			}
			if tt.wantContain != "" && !strings.Contains(w.Body.String(), tt.wantContain) {
				t.Errorf("body = %q, want to contain %q", w.Body.String(), tt.wantContain)
			}
		})
	}
}

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name        string
		form        url.Values
		wantStatus  int
		wantContain string
	}{
		{
			name:        "valid input returns probability",
			form:        url.Values{"p": {"0.5"}, "n": {"4"}, "k": {"2"}},
			wantStatus:  http.StatusOK,
			wantContain: "0.375",
		},
		{
			name:        "k exceeds n returns error containing 'cannot'",
			form:        url.Values{"p": {"0.5"}, "n": {"3"}, "k": {"5"}},
			wantStatus:  http.StatusOK,
			wantContain: "cannot",
		},
		{
			name:        "invalid p re-renders form with error",
			form:        url.Values{"p": {"abc"}, "n": {"4"}, "k": {"2"}},
			wantStatus:  http.StatusOK,
			wantContain: "<form",
		},
		{
			name:        "invalid n re-renders form with error",
			form:        url.Values{"p": {"0.5"}, "n": {"abc"}, "k": {"2"}},
			wantStatus:  http.StatusOK,
			wantContain: "<form",
		},
		{
			name:        "invalid k re-renders form with error",
			form:        url.Values{"p": {"0.5"}, "n": {"4"}, "k": {"abc"}},
			wantStatus:  http.StatusOK,
			wantContain: "<form",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/calculate", strings.NewReader(tt.form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			calculateHandler(w, req)
			if w.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.wantStatus)
			}
			if !strings.Contains(w.Body.String(), tt.wantContain) {
				t.Errorf("body = %q, want to contain %q", w.Body.String(), tt.wantContain)
			}
		})
	}
}
