package mail

import (
	"strings"
	"testing"
)

func TestHTMLMessageBodyEscapesCode(t *testing.T) {
	body := htmlMessageBody(`<script>alert("x")</script>`)

	if strings.Contains(body, "<script>") {
		t.Fatalf("html body contains unescaped code: %s", body)
	}

	if !strings.Contains(body, "&lt;script&gt;") {
		t.Fatalf("html body does not contain escaped code: %s", body)
	}
}

func TestPlainMessageBodyContainsCode(t *testing.T) {
	const code = "123456"

	body := plainMessageBody(code)
	if !strings.Contains(body, code) {
		t.Fatalf("plain body does not contain code: %s", body)
	}
}
