package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestFormHandler_GET_renders_form(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	formHandler(w, req)
	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Result().StatusCode)
	}
	if !strings.Contains(w.Body.String(), "Binomial Probability Calculator") {
		t.Error("expected title in body")
	}
}

func TestCalculateHandler_valid_input_shows_result_and_preserves_form(t *testing.T) {
	form := url.Values{"p": {"0.5"}, "n": {"10"}, "k": {"3"}}
	req := httptest.NewRequest(http.MethodPost, "/calculate", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	calculateHandler(w, req)
	body := w.Body.String()
	if !strings.Contains(body, "0.1171875") {
		t.Errorf("expected result in body, got:\n%s", body)
	}
	if !strings.Contains(body, `value="0.5"`) {
		t.Errorf("expected p pre-filled, got:\n%s", body)
	}
	if !strings.Contains(body, `value="10"`) {
		t.Errorf("expected n pre-filled, got:\n%s", body)
	}
	if !strings.Contains(body, `value="3"`) {
		t.Errorf("expected k pre-filled, got:\n%s", body)
	}
}

func TestCalculateHandler_invalid_p_shows_error_and_preserves_form(t *testing.T) {
	form := url.Values{"p": {"abc"}, "n": {"10"}, "k": {"3"}}
	req := httptest.NewRequest(http.MethodPost, "/calculate", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	calculateHandler(w, req)
	body := w.Body.String()
	if !strings.Contains(body, "Invalid value for p") {
		t.Errorf("expected error message, got:\n%s", body)
	}
	if !strings.Contains(body, `value="abc"`) {
		t.Errorf("expected p pre-filled on error, got:\n%s", body)
	}
	if !strings.Contains(body, `value="10"`) {
		t.Errorf("expected n pre-filled on error, got:\n%s", body)
	}
}

func TestCalculateHandler_calc_error_shows_error_and_preserves_form(t *testing.T) {
	// k > n is invalid for binomial probability
	form := url.Values{"p": {"0.5"}, "n": {"3"}, "k": {"10"}}
	req := httptest.NewRequest(http.MethodPost, "/calculate", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	calculateHandler(w, req)
	body := w.Body.String()
	if !strings.Contains(body, `value="0.5"`) {
		t.Errorf("expected p pre-filled on calc error, got:\n%s", body)
	}
	if !strings.Contains(body, "class=\"error\"") {
		t.Errorf("expected error message div, got:\n%s", body)
	}
	// no result card should appear
	if strings.Contains(body, "result-value") {
		t.Errorf("expected no result card on error, got:\n%s", body)
	}
}
