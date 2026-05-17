package res

import (
	"net/http"
	"strings"
	"testing"
)

func TestDecodeBodyRejectsUnknownFields(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"ring","extra":true}`))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	var payload struct {
		Name string `json:"name"`
	}
	if err := DecodeBody(req, &payload); err == nil {
		t.Fatal("expected unknown field to be rejected")
	}
}

func TestDecodeBodyRejectsMultipleJSONValues(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"ring"} {"name":"bracelet"}`))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	var payload struct {
		Name string `json:"name"`
	}
	if err := DecodeBody(req, &payload); err == nil {
		t.Fatal("expected multiple JSON values to be rejected")
	}
}
