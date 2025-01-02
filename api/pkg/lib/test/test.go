package tu

import (
	"bytes"
	"io"
	"log/slog"
	"mime/multipart"
)

func NewTestLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}

func NewTestFile(content string) multipart.File {
	fileBuffer := bytes.NewBufferString(content)

	return &mockMultipartFile{
		Reader: bytes.NewReader(fileBuffer.Bytes()),
	}
}

type mockMultipartFile struct {
	*bytes.Reader
}

func (m *mockMultipartFile) Close() error {
	return nil
}
