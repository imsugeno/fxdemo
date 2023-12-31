package main

import (
	"io"
	"net/http"

	"go.uber.org/zap"
)

// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct {
	logger *zap.Logger
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(logger *zap.Logger) *EchoHandler {
	return &EchoHandler{
		logger: logger,
	}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.logger.Warn("Failed to handle request:", zap.Error(err))
	}
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}
